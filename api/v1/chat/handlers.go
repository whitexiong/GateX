package chat

import (
	"gateway/api/v1/setting"
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
				IsAI:       dbUser.IsAI,
			})
		} else {
			users = append(users, models.ChatUserResponse{
				ID:         dbUser.ID,
				Username:   dbUser.Username,
				Avatar:     dbUser.AvatarUrl,
				ChatRoomID: chatRoomUser.ChatRoomID,
				IsAI:       dbUser.IsAI,
			})
		}
	}

	setting.SendResponse(c, http.StatusOK, 200, users)
}

func GetChatHistoryForRoom(c *gin.Context) {
	currentUserIdStr, ok := c.Get("user_id")
	if !ok {
		setting.SendResponse(c, http.StatusBadRequest, 400, "User ID not found or not of expected type.")
		return
	}

	// 这里简化了用户ID的处理，并假设它是uint类型
	currentUserId := uint(currentUserIdStr.(float64))

	chatRoomId, err := strconv.Atoi(c.Param("chatRoomId"))
	if err != nil {
		setting.SendResponse(c, http.StatusBadRequest, 400, "Invalid chat room ID.")
		return
	}

	var messages []models.Message
	if err := models.DB.Where("chat_room_id = ?", chatRoomId).
		Order("created_at asc").Find(&messages).Error; err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, 500, "Error fetching chat history.")
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

	setting.SendResponse(c, http.StatusOK, 200, response)
}

func GetChatWindowsByUser(c *gin.Context) {
	currentUserIdStr, ok := c.Get("user_id")
	if !ok {
		setting.SendResponse(c, http.StatusBadRequest, 400, "User ID not found or not of expected type.")
		return
	}
	currentUserId := uint(currentUserIdStr.(float64))

	var chatRooms []models.ChatRoom

	result := models.DB.Preload("Users").Preload("Messages").Joins("JOIN chat_room_users on chat_room_users.chat_room_id = chat_rooms.id").Where("chat_room_users.user_id = ?", currentUserId).Find(&chatRooms)
	if result.Error != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to retrieve chat rooms.")
		return
	}

	setting.SendResponse(c, http.StatusOK, 200, chatRooms)
}

func CreateChatRoom(c *gin.Context) {
	var request models.ChatRoomRequest

	if err := c.BindJSON(&request); err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid input data.")
		return
	}

	currentUserIdStr, ok := c.Get("user_id")
	if !ok {
		setting.SendResponse(c, http.StatusBadRequest, 400, "User ID not found or not of expected type.")
		return
	}

	currentUserId := uint(currentUserIdStr.(float64))
	request.UserIDs = append(request.UserIDs, currentUserId) // 加入到UserIDs中

	var users []*models.User
	if err := models.DB.Where("id IN ?", request.UserIDs).Find(&users).Error; err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to find users.")
		return
	}

	hasAIUser := false
	for _, user := range users {
		if user.IsAI {
			hasAIUser = true
			break
		}
	}

	if hasAIUser {
		request.RoomType = models.AIChatRoom
	}

	roomName := ""
	for i, user := range users {
		roomName += user.Username
		if i != len(users)-1 {
			roomName += "、"
		}
	}

	chatRoom := models.ChatRoom{
		Name:        roomName,
		Description: request.Description,
		RoomType:    request.RoomType,
		Users:       users,
	}

	tx := models.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&chatRoom).Error; err != nil {
		tx.Rollback()
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to create chat room.")
		return
	}

	if err := tx.Commit().Error; err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to commit transaction.")
		return
	}

	setting.SendResponse(c, http.StatusOK, 200, nil)
}

func DeleteChatWindow(c *gin.Context) {
	roomId := c.Param("roomId")
	var chatRoom models.ChatRoom

	result := models.DB.Where("id = ?", roomId).Delete(&chatRoom)
	if result.Error != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to delete chat room.")
		return
	}

	setting.SendResponse(c, http.StatusOK, 200, "Chat room deleted successfully.")
}
