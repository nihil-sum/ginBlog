package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/teris-io/shortid"
)

func GenShortID() (string, error) {
	return shortid.Generate()
}

func Md5(str string) (string, error) {
	h := md5.New()

	_, err := io.WriteString(h, str)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func RandStr(n int) string {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	const pattern = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	salt := make([]byte, n)
	l := len(pattern)

	for i := 0; i < n; i++ {
		p := r.Intn(l)
		salt = append(salt, pattern[p])

	}
	return string(salt)
}

func GetHostName() string {
	name, err := os.Hostname()
	if err != nil {
		return ""
	}
	return name
}

func GetLastTenDayTimestamp() string {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).AddDate(0, 0, -10).Format("2006-01-02 15:04:05")
	return startTime
}
