package auth

import (
	"fmt"
	"gateway/handlers"
	"gateway/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var jwtService = JWT{
	SecretKey: os.Getenv("JWT_SECRET_KEY"),
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

	roleNames := extractRoleNames(user.Roles)

	var primaryRole string
	if len(roleNames) > 0 {
		primaryRole = roleNames[0]
	} else {
		primaryRole = ""
	}

	token, err := jwtService.GenerateToken(user.Username, int64(user.ID), primaryRole)
	if err != nil {
		panic("CustomError#500#Could not generate token")
	}

	handlers.SendResponse(c, http.StatusOK, 200, gin.H{
		"token": token,
		"user": gin.H{
			"username": user.Username,
			"id":       user.ID,
			"roles":    roleNames,
		},
	}, "Success")
	return
}

func extractRoleNames(roles []*models.Role) []string {
	var roleNames []string
	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}
	return roleNames
}

func Logout(c *gin.Context) {
	handlers.SendResponse(c, http.StatusOK, 200, nil, "Success")
	return
}
