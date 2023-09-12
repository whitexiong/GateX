package handlers

import (
	"gateway/apierrors"
	"gateway/models"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	if result := models.DB.Preload("Roles").Find(&users); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
		return
	}
	SendResponse(c, http.StatusOK, 200, users)
}

func CreateUser(c *gin.Context) {
	var userReq models.UserRequest

	if err := c.BindJSON(&userReq); err != nil {
		SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
		return
	}

	userReq.Password = string(hashedPassword)

	if result := models.DB.Create(&userReq.User); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
		return
	}

	for _, roleID := range userReq.Roles {
		userRole := models.UserRole{
			UserID:    userReq.ID,
			RoleID:    roleID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if result := models.DB.Create(&userRole); result.Error != nil {
			SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
			return
		}
	}

	SendResponse(c, http.StatusOK, 200, userReq.User)
}

func GetUserDetail(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if result := models.DB.Preload("Roles").First(&user, id); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
		return
	}

	responseData := structs.Map(user)

	roleIDs := make([]int, len(user.Roles))
	for i, role := range user.Roles {
		roleIDs[i] = int(role.ID)
	}
	responseData["Roles"] = roleIDs

	SendResponse(c, http.StatusOK, 200, responseData)
}

func UpdateUser(c *gin.Context) {
	var userReq models.UserRequest
	id := c.Param("id")

	if result := models.DB.First(&userReq.User, id); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
		return
	}

	if err := c.BindJSON(&userReq); err != nil {
		SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
		return
	}

	if result := models.DB.Save(&userReq.User); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
		return
	}

	models.DB.Where("user_id = ?", userReq.ID).Delete(models.UserRole{})
	for _, roleID := range userReq.Roles {
		userRole := models.UserRole{
			UserID:    userReq.ID,
			RoleID:    roleID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		models.DB.Create(&userRole)
	}

	SendResponse(c, http.StatusOK, 200, userReq.User)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	models.DB.Where("user_id = ?", id).Delete(models.UserRole{})

	if result := models.DB.Delete(&user, id); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
		return
	}

	SendResponse(c, http.StatusOK, 200, nil)
}
