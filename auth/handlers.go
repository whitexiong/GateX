package auth

import (
	"fmt"
	"gateway/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var jwtService = JWT{
	SecretKey: "YOUR_SECRET_KEY", // 这里需要更安全的键
	Issuer:    "gatewayApp",
	Expiry:    24 * time.Hour,
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		panic("CustomError#400#" + err.Error())
	}

	// 从数据库获取用户
	user, err := models.FindUserByUsername(input.Username)
	if err != nil {
		panic(fmt.Sprintf("CustomError#400# %v", err))
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		panic("CustomError#401#Invalid password")
	}

	token, err := jwtService.GenerateToken(input.Username)
	if err != nil {
		panic("CustomError#500#Could not generate token")
	}

	c.JSON(200, gin.H{
		"token": token,
		"user": gin.H{
			"username": user.Username,
		},
	})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "退出登录成功！",
	})
}
