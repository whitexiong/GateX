package main

import (
	"fmt"
	"gateway/api/v1/handlers"
	"gateway/api/v1/routes"
	"gateway/middleware"
	"gateway/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("加载.env文件错误: %s", err)
		return
	}

	// 从环境变量中读取服务器的端口
	proxyPort := os.Getenv("SERVER_PROXY_PORT")
	panelPort := os.Getenv("SERVER_PANEL_PORT")

	// 1. 初始化数据库
	err = models.InitDatabase()
	if err != nil {
		log.Fatalf("连接数据库失败 %v", err)
	}

	// 2. 请求代理
	go func() {
		http.HandleFunc("/", handlers.Proxy)
		log.Printf("Server started on :%s", proxyPort)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", proxyPort), nil))
	}()

	// 3. 启动面板
	r := gin.Default()

	// 初始化中间件
	middlewares := middleware.InitializeMiddlewares()
	for _, m := range middlewares {
		r.Use(m)
	}

	// 启动路由
	for _, setupFunc := range routes.AllRoutes {
		setupFunc(r)
	}

	// 其他路由和中间件设置
	err = r.Run(fmt.Sprintf(":%s", panelPort))
	if err != nil {
		return
	}
}
