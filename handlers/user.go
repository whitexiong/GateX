package handlers

import (
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		SendResponse(c, http.StatusBadRequest, 400, nil, "Invalid request data")
		return
	}

	if result := models.DB.Create(&user); result.Error != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to create user")
		return
	}

	SendResponse(c, http.StatusOK, 200, user, "Success")
}

func GetUserDetail(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if result := models.DB.Preload("Roles").First(&user, id); result.Error != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to fetch user")
		return
	}

	SendResponse(c, http.StatusOK, 200, user, "Success")
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if result := models.DB.First(&user, id); result.Error != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to fetch user")
		return
	}

	if err := c.BindJSON(&user); err != nil {
		SendResponse(c, http.StatusBadRequest, 400, nil, "Invalid request data")
		return
	}

	if result := models.DB.Save(&user); result.Error != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to update user")
		return
	}

	SendResponse(c, http.StatusOK, 200, user, "Success")
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if result := models.DB.Delete(&user, id); result.Error != nil {
		SendResponse(c, http.StatusInternalServerError, 500, nil, "Failed to delete user")
		return
	}

	SendResponse(c, http.StatusOK, 200, nil, "User deleted successfully")
}
