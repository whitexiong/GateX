package main

import (
	"gateway/handlers"
	"gateway/middleware"
	"gateway/models"
	"gateway/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// 1. 初始化数据库
	err := models.InitDatabase()
	if err != nil {
		log.Fatalf("连接数据库失败 %v", err)
	}

	// 2. 请求代理
	go func() { // 使用 go 关键字启动新的 goroutine
		http.HandleFunc("/", handlers.Proxy)
		log.Println("Server started on :8050")
		log.Fatal(http.ListenAndServe(":8050", nil))
	}()

	// 3. 启动面板
	r := gin.Default()

	config := cors.DefaultConfig()

	// 明确允许的来源
	config.AllowOrigins = []string{"http://192.168.0.210:8080"}

	// 允许的 HTTP 方法
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	// 允许的请求头。注意我添加了"Authorization"，以处理JWT或其他认证方法。
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	// 允许客户端携带验证信息，例如cookie。如果你使用基于cookie的身份验证，这将是必需的。
	config.AllowCredentials = true

	// 使用CORS中间件
	r.Use(cors.New(config))

	// 初始化中间件
	middlewares := middleware.InitializeMiddlewares()
	for _, m := range middlewares {
		r.Use(m)
	}

	//启动路由
	for _, setupFunc := range routes.AllRoutes {
		setupFunc(r)
	}

	// 其他路由和中间件设置
	err = r.Run(":8051")
	if err != nil {
		return
	}
}
