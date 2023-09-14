package dao

import "gateway/models"

func GetChatUsers() ([]models.User, error) {
	var users []models.User
	// Fetch users based on your logic. For example, all users, online users, etc.
	if err := models.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
