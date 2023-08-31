package middleware

import (
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	secretKey = "YOUR_SECRET_KEY" // 你的JWT签名密钥，实际项目中不应硬编码此值，而应该从环境变量或配置文件中读取。
)

func InitJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取令牌

		excludedPaths := []string{"/user/login", "/user/logout"}
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

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header not provided"})
			return
		}

		// 接受 Bearer Token 的格式，所以我们需要将它分割开来
		splitToken := strings.Split(tokenString, "Bearer ")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
			return
		}
		tokenString = splitToken[1]

		// 验证令牌
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return []byte(secretKey), nil
		})

		// 处理令牌验证的结果
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token", "error": err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("username", claims["username"])
			c.Set("role", claims["role"])
			c.Set("user_id", claims["user_id"])
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token11111"})
	}
}
