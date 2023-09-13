package auth

import (
	"gateway/api/v1/handlers"
	"gateway/apierrors"
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
		handlers.SendResponse(c, http.StatusOK, apierrors.InvalidRequestData, nil)
		return
	}

	user, err := models.FindUserByUsername(input.Username)
	if err != nil {
		handlers.SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		handlers.SendResponse(c, http.StatusOK, apierrors.AuthenticationFailed, nil)
		return
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
		c.Error(err)
	}

	handlers.SendResponse(c, http.StatusOK, 200, gin.H{
		"token": token,
		"user": gin.H{
			"username":   user.Username,
			"avatar_url": user.AvatarUrl,
			"id":         user.ID,
			"roles":      roleNames,
		},
	})
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
	handlers.SendResponse(c, http.StatusOK, 200, nil)
	return
}
