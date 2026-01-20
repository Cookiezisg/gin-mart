package main

import (
	"fmt"
	"golang-mvc-ecommerce/api"
	"golang-mvc-ecommerce/utils/graceful"
	"log"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           电商项目 API 接口文档
// @version         1.0
// @description     这是一个基于 Gin 框架开发的电商系统后端接口。
// @termsOfService  http://swagger.io/terms/

// @contact.name   无敌我最帅

// @host      localhost:8080
// @BasePath  /
func main() {
	r := gin.Default()
	registerMiddlewares(r)
	api.RegisterHandlers(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\\n", err)
		}
	}()
	graceful.ShutdownGin(srv, time.Second*5)

}

// 注册中间件
func registerMiddlewares(r *gin.Engine) {
	r.Use(
		gin.LoggerWithFormatter(
			func(param gin.LogFormatterParams) string {

				return fmt.Sprintf(
					"%s - [%s] \"%s %s %s\" %d %s \"%s\"\n",
					param.ClientIP,
					param.TimeStamp.Format(time.RFC3339),
					param.Method,
					param.Path,
					param.Request.Proto,
					param.StatusCode,
					param.Latency,
					param.ErrorMessage,
				)
			}))
	r.Use(gin.Recovery())
}
