package middleware

import (
	"gateway/models"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func InitializeCasbinMiddleware() gin.HandlerFunc {
	// 使用GORM适配器
	adapter, err := gormadapter.NewAdapterByDB(models.DB)

	if err != nil {
		log.Fatalf("Failed to initialize the GORM adapter: %v", err)
		return nil
	}

	enforcer, err := casbin.NewEnforcer("config/gateway.conf", adapter)

	if err != nil {
		log.Fatalf("Failed to initialize casbin with GORM adapter: %v", err)
		return nil
	}

	return func(c *gin.Context) {

		excludedPaths := []string{"/user/login", "/user/logout", "/uploads/", "/ws"}
		currentPath := c.FullPath()
		log.Printf("Current request path: %s", currentPath)

		for _, path := range excludedPaths {
			log.Printf("Checking path: %s", path)
			if currentPath == path {
				log.Println("Path matched excluded path. Skipping Casbin check.")
				c.Next()
				return
			}
		}

		roleValue, exists := c.Get("role")
		if !exists {
			log.Println("Role not found in context. Ensure JWT middleware is run before Casbin middleware.")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Role not found",
			})
			return
		}

		role, ok := roleValue.(string)
		if !ok {
			log.Println("Role in context is not of type string.")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server apierrors",
			})
			return
		}

		sub := role

		obj := c.Request.URL.Path
		act := c.Request.Method

		log.Println("sub, obj, act", sub, obj, act)
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
