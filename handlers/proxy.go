package handlers

import (
	"gateway/models"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

var target, _ = url.Parse("http://fire.com:80")
var proxy = httputil.NewSingleHostReverseProxy(target)

func Proxy(w http.ResponseWriter, r *http.Request) {
	// 1. 解析请求数据
	startTime := time.Now() // 记录请求开始时间
	log.Printf("Received request %s %s", r.Method, r.URL.String())
	path := r.URL.Path

	// 2. 判断是否为新接口
	var api models.APIEndpoint
	result := models.DB.Where("path = ?", path).First(&api)
	if result.Error != nil { // 如果接口不在数据库中
		api.Path = path
		api.HitCount = 1
		api.Type = detectResourceType(path) // 3. 区分请求的资源类型
		api.RequestMethod = r.Method        // 记录请求方法
		models.DB.Create(&api)
	} else { // 如果接口已经在数据库中
		api.HitCount++
		api.RequestMethod = r.Method
		models.DB.Save(&api)
	}

	proxy.ServeHTTP(w, r)

	// 计算请求时间
	requestDuration := time.Since(startTime).Milliseconds()

	requestDurationUint64 := uint64(requestDuration)

	if api.MinRequestTime == 0 || requestDurationUint64 < api.MinRequestTime {
		api.MinRequestTime = requestDurationUint64
	}
	if requestDurationUint64 > api.MaxRequestTime {
		api.MaxRequestTime = requestDurationUint64
	}

	models.DB.Save(&api)
}

func detectResourceType(path string) string {
	if strings.HasSuffix(path, ".js") {
		return "JavaScript"
	} else if strings.HasSuffix(path, ".css") {
		return "CSS"
	} else if strings.HasSuffix(path, ".png") || strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") {
		return "Image"
	} else if strings.HasSuffix(path, ".ico") {
		return "Icon"
	} else {
		return "API" // 默认为 API 类型
	}
}
