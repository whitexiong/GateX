package handlers

import (
	"fmt"
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
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to fetch users")
		return
	}
	SendResponse(c, http.StatusOK, 200, users, "Success")
}

func CreateUser(c *gin.Context) {
	var userReq models.UserRequest

	if err := c.BindJSON(&userReq); err != nil {
		SendResponse(c, http.StatusBadRequest, 400, nil, "Invalid request data")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Password hashing failed")
		return
	}

	userReq.Password = string(hashedPassword)

	if result := models.DB.Create(&userReq.User); result.Error != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to create user")
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
			SendResponse(c, http.StatusInternalServerError, 500, nil, fmt.Sprintf("Failed to create user-role relation for Role ID %d", roleID))
			return
		}
	}

	SendResponse(c, http.StatusOK, 200, userReq.User, "Success")
}

func GetUserDetail(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if result := models.DB.Preload("Roles").First(&user, id); result.Error != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to fetch user")
		return
	}

	responseData := structs.Map(user)

	roleIDs := make([]int, len(user.Roles))
	for i, role := range user.Roles {
		roleIDs[i] = int(role.ID)
	}
	responseData["Roles"] = roleIDs

	SendResponse(c, http.StatusOK, 200, responseData, "Success")
}

func UpdateUser(c *gin.Context) {
	var userReq models.UserRequest
	id := c.Param("id")

	if result := models.DB.First(&userReq.User, id); result.Error != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to fetch user")
		return
	}

	if err := c.BindJSON(&userReq); err != nil {
		SendResponse(c, http.StatusBadRequest, 400, nil, "Invalid request data")
		return
	}

	// 更新用户的基本属性
	if result := models.DB.Save(&userReq.User); result.Error != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to update user")
		return
	}

	// 更新用户的角色关系：先删除现有关系，再添加新的关系
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

	SendResponse(c, http.StatusOK, 200, userReq.User, "Success")
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	// 删除与用户相关的所有角色关系
	models.DB.Where("user_id = ?", id).Delete(models.UserRole{})

	// 删除用户本身
	if result := models.DB.Delete(&user, id); result.Error != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to delete user")
		return
	}

	SendResponse(c, http.StatusOK, 200, nil, "User deleted successfully")
}
