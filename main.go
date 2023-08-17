package main

import (
	"gateway/handlers"
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

	//err = middleware.InitializeCasbin()
	//if err != nil {
	//	log.Fatalf("Failed to initialize casbin: %v", err)
	//}

	// 3. 启动面板
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://192.168.0.210:8080"}    // 可以设置为 "*" 允许所有来源，但出于安全性考虑最好明确指定来源
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"} // 允许的 HTTP 方法
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}

	r.Use(cors.New(config))

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
