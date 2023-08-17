package middleware

import (
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

var enforcer *casbin.Enforcer

func InitializeCasbin() error {
	var err error
	enforcer, err = casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		return err
	}
	return nil
}

func Authorize(obj string, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := c.GetString("username") // 假设你从上下文或其他方式获取用户名

		// Check if the user.go has the permission to access the object-action
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

		c.Next()
	}
}
