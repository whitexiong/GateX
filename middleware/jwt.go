package middleware

import (
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var secretKey = os.Getenv("JWT_SECRET_KEY")

func InitJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		excludedPaths := []string{"/setting/user/login", "/setting/user/logout", "/uploads/", "/ws"}
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

		splitToken := strings.Split(tokenString, "Bearer ")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
			return
		}
		tokenString = splitToken[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token", "apierrors": err.Error()})
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
