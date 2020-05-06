package router

import (
	"fmt"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func InitRoutes() *gin.Engine {
	engine := gin.New()
	// 使用zap日志库
	engine.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(zap.L(), true))

	g := engine.Group("/api/v1")
	{
		// Example ping request.
		g.GET("/hello", func(c *gin.Context) {
			c.String(200, "hello world !!! "+fmt.Sprint(time.Now().Format("2006-01-2 15:04:05")))
		})

		// Example when panic happen.
		g.GET("/panic", func(c *gin.Context) {
			panic("An unexpected error happen!")
		})
	}

	return engine
}
