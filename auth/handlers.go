package auth

import (
	"fmt"
	"gateway/handlers"
	"gateway/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

	// 从user.Roles字段中提取角色名
	roleNames := extractRoleNames(user.Roles)

	// 传递username, userID, and roles到GenerateToken函数
	// 这里我假设你希望在token中只包含第一个角色名，如果不是这样，请根据需要修改。
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
