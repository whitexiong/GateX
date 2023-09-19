package chat

import (
	"gateway/api/v1/handlers"
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetChatUserList(c *gin.Context) {
	var users []models.ChatUserResponse
	var dbUsers []models.User

	models.DB.Find(&dbUsers)

	for _, dbUser := range dbUsers {
		var chatRoomUser models.ChatRoomUser

		result := models.DB.Where("user_id = ?", dbUser.ID).First(&chatRoomUser)
		if result.Error != nil {
			users = append(users, models.ChatUserResponse{
				ID:         dbUser.ID,
				Username:   dbUser.Username,
				Avatar:     dbUser.AvatarUrl,
				ChatRoomID: 0,
			})
		} else {
			users = append(users, models.ChatUserResponse{
				ID:         dbUser.ID,
				Username:   dbUser.Username,
				Avatar:     dbUser.AvatarUrl,
				ChatRoomID: chatRoomUser.ChatRoomID,
			})
		}
	}

	handlers.SendResponse(c, http.StatusOK, 200, users)
}

func GetChatHistoryForUser(c *gin.Context) {
	currentUserIdStr, ok := c.Get("user_id")
	if !ok {
		handlers.SendResponse(c, http.StatusBadRequest, 400, "User ID not found or not of expected type.")
		return
	}

	// 这里简化了用户ID的处理，并假设它是uint类型
	currentUserId := uint(currentUserIdStr.(float64))

	otherUserId, err := strconv.Atoi(c.Param("otherUserId"))
	if err != nil {
		handlers.SendResponse(c, http.StatusBadRequest, 400, "Invalid other user ID.")
		return
	}

	var messages []models.Message
	if err := models.DB.Where("(sender_id = ? AND to_user_id = ?) OR (sender_id = ? AND to_user_id = ?)",
		currentUserId, otherUserId, otherUserId, currentUserId).
		Order("created_at asc").Find(&messages).Error; err != nil {
		handlers.SendResponse(c, http.StatusInternalServerError, 500, "Error fetching chat history.")
		return
	}

	type ChatHistoryResponse struct {
		Content string `json:"content"`
		Type    string `json:"type"`
	}

	var response []ChatHistoryResponse
	for _, msg := range messages {
		msgType := "received"
		if msg.SenderID == currentUserId {
			msgType = "sent"
		}
		response = append(response, ChatHistoryResponse{
			Content: msg.Content,
			Type:    msgType,
		})
	}

	handlers.SendResponse(c, http.StatusOK, 200, response)
}

func GetChatWindow(c *gin.Context) {
	roomId := c.Param("roomId")
	var chatRoom models.ChatRoom

	result := models.DB.Preload("Users").Preload("Messages").Where("id = ?", roomId).First(&chatRoom)
	if result.Error != nil {
		handlers.SendResponse(c, http.StatusInternalServerError, -1, "Failed to retrieve chat room.")
		return
	}

	handlers.SendResponse(c, http.StatusOK, 200, chatRoom)
}

func CreateChatRoom(c *gin.Context) {
	var chatRoom models.ChatRoom

	if err := c.BindJSON(&chatRoom); err != nil {
		handlers.SendResponse(c, http.StatusBadRequest, -1, "Invalid input data.")
		return
	}

	result := models.DB.Create(&chatRoom)
	if result.Error != nil {
		handlers.SendResponse(c, http.StatusInternalServerError, -1, "Failed to create chat room.")
		return
	}

	handlers.SendResponse(c, http.StatusOK, 200, chatRoom)
}

func DeleteChatWindow(c *gin.Context) {
	roomId := c.Param("roomId")
	var chatRoom models.ChatRoom

	result := models.DB.Where("id = ?", roomId).Delete(&chatRoom)
	if result.Error != nil {
		handlers.SendResponse(c, http.StatusInternalServerError, -1, "Failed to delete chat room.")
		return
	}

	handlers.SendResponse(c, http.StatusOK, 200, "Chat room deleted successfully.")
}
