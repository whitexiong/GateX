package middleware

import (
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func InitializeCasbinMiddleware() gin.HandlerFunc {
	// 这里，我们在函数内部初始化了 enforcer
	enforcer, err := casbin.NewEnforcer("config/model.conf", "config/policy.cvs")
	if err != nil {
		log.Fatalf("Failed to initialize casbin: %v", err)
		return nil
	}

	return func(c *gin.Context) {
		// 获取权限相关的信息
		//sub := c.GetString("username")  // 这里假设"username"是存储在Gin的上下文中的当前用户的键。
		sub := "alice"            // 这里假设"username"是存储在Gin的上下文中的当前用户的键。
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
