package middleware

import (
	"gateway/models"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitializeCasbinMiddleware() gin.HandlerFunc { // 假设你已经有一个 GORM 的 db 连接
	// 使用GORM适配器
	adapter, err := gormadapter.NewAdapterByDB(models.DB)
	if err != nil {
		log.Fatalf("Failed to initialize the GORM adapter: %v", err)
		return nil
	}

	enforcer, err := casbin.NewEnforcer("config/model.conf", adapter)
	if err != nil {
		log.Fatalf("Failed to initialize casbin with GORM adapter: %v", err)
		return nil
	}

	return func(c *gin.Context) {
		// 获取权限相关的信息
		sub := "alice"            // 这里假设"username"是存储在Gin的上下文中的当前用户的键。 这里是为了测试正确性
		obj := c.Request.URL.Path // 示例：获取请求的URL作为对象
		act := c.Request.Method   // 示例：获取HTTP方法作为动作

		// 检查用户是否具有权限访问 object-action
		ok, err := enforcer.Enforce(sub, obj, act)
		if err != nil {
			log.Println("Error occurred while enforcing:", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Error occurred while enforcing",
			})
			return
		}

		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Permission denied",
			})
			return
		}

		c.Next() // 确保请求传递给下一个中间件/处理程序
	}
}
