package main

import (
	"gin-example/log"
	"gin-example/router"
	"go.uber.org/zap"
)

func init() {
	// 初始化日志库
	log.SetLogs(zap.DebugLevel, log.LOGFORMAT_CONSOLE, "./log/gin-example.log")
}

func main() {
	r := router.InitRoutes()

	//启动监听
	if err := r.Run(":8888"); err != nil {
		zap.L().Fatal("HTTP Server启动失败", zap.Error(err))
	}
}
