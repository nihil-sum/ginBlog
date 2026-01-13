package common

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"ginBlog/conf"
	"github.com/satori/go.uuid"
	"gopkg.in/vansante/go-ffprobe.v2"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func getVideoInfo(path string) (*ffprobe.ProbeData, error) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelFn()
	ffprobePath := conf.Conf.App.Ffprobepath
	ffprobe.SetFFProbeBinPath(ffprobePath)
	data, err := ffprobe.ProbeURL(ctx, path)

	if err != nil {
		return nil, err
	}

	return data, err
}

func DownloadFile(uri string) (string, error) {
	fileName := uuid.NewV4().String() + filepath.Base(uri)
	path := GetTmpFilePath() + "/" + fileName

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
	if err != nil {
		return "", nil
	}
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Address not found")
	}
	defer resp.Body.Close()

	out, err := os.Create(path)

	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	return path, nil

}
func GetFileType(format string) string {
	if StringContains(format, "jpg", "jpeg", "png", "png") {
		return "image"
	} else if StringContains(format, "avi", "mp4", "wmv", "mpg", "mov", "mpeg", "rm", "swf", "flv", "ram") {
		return "video"
	} else {
		return format
	}
}

func GetFileMd5(filename string) (string, error) {
	pFile, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("Failed to open the file")
		return "", err
	}
	defer pFile.Close()

	md5h := md5.New()
	_, err = io.Copy(md5h, pFile)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5h.Sum(nil)), nil
}

func GetFilePath() string {
	folderName := time.Now().Format("20251217")
	folderPath := filepath.Join(conf.Conf.App.Resource, "/", folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, os.ModePerm)
	}
	return folderPath
}

func GetTmpFilePath() string {
	folderName := time.Now.Format("20251217")
	folderPath := filepath.Join(conf.Conf.App.Resource, "/tmp", folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, os.ModePerm)
	}

	return folderPath
}
