package main

import (
	"context"
	"ginBlog/conf"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"log"
	"net/http"
	"time"
)

var (
	cfgFile = pflag.StringP("config", "c", "./conf/config.yml", "config file path")
	version = pflag.BoolP("version", "v", false, "show version info")
)

func main() {
	pflag.Parse()
	if *version {
		log.Println("Version", "v2.0")
	}
	cfg := conf.Init(*cfgFile)
	logger.Init(&cfg.Logger)
	repository.Init(&cfg.ORM)
	gin.SetMode(cfg.App.Mode)

	log.Println("HTTP服务器启动",cfg.App.Addr)
	logger.Info("HTTP服务器启动")

	srv := &http.Server{
		Addr: cfg.App.Addr
		Handler
		
	}

}
