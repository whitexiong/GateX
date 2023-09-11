package handlers

import (
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
		c.Error(result.Error)
		return
	}
	SendResponse(c, http.StatusOK, 200, users)
}

func CreateUser(c *gin.Context) {
	var userReq models.UserRequest

	if err := c.BindJSON(&userReq); err != nil {
		c.Error(err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Error(err)
		return
	}

	userReq.Password = string(hashedPassword)

	if result := models.DB.Create(&userReq.User); result.Error != nil {
		c.Error(err)
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
			c.Error(err)
			return
		}
	}

	SendResponse(c, http.StatusOK, 200, userReq.User)
}

func GetUserDetail(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if result := models.DB.Preload("Roles").First(&user, id); result.Error != nil {
		c.Error(result.Error)
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
		c.Error(result.Error)
		return
	}

	if err := c.BindJSON(&userReq); err != nil {
		c.Error(err)
		return
	}

	// 更新用户的基本属性
	if result := models.DB.Save(&userReq.User); result.Error != nil {
		c.Error(result.Error)
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

	SendResponse(c, http.StatusOK, 200, userReq.User)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	// 删除与用户相关的所有角色关系
	models.DB.Where("user_id = ?", id).Delete(models.UserRole{})

	// 删除用户本身
	if result := models.DB.Delete(&user, id); result.Error != nil {
		c.Error(result.Error)
		return
	}

	SendResponse(c, http.StatusOK, 200, nil)
}
