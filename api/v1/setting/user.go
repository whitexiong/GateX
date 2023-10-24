package setting

import (
	"gateway/apierrors"
	"gateway/models"
	"gateway/pkg/pagination"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

func GetAllUsers(c *gin.Context) {
	type RequestBody struct {
		Name        string `json:"name"`
		PageSize    int    `json:"pageSize"`
		CurrentPage int    `json:"currentPage"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		SendResponse(c, http.StatusBadRequest, apierrors.InvalidRequestData, nil)
		return
	}

	paginator, searcher := pagination.NewPaginationService(body.CurrentPage, body.PageSize, body.Name,
		pagination.WithSearchField("Username"),
	)

	query := models.DB.Model(&models.User{}).Preload("Roles").Where("Username != ?", "admin")
	query = searcher.Search(query)
	query = paginator.Paginate(query)

	var users []models.User
	if result := query.Find(&users); result.Error != nil {
		SendResponse(c, http.StatusInternalServerError, apierrors.DataNotFound, nil)
		return
	}

	columnsConfig := []map[string]interface{}{
		{
			"label":      "用户名",
			"key":        "Username",
			"sortable":   true,
			"searchable": true,
		},
		{
			"label":      "邮箱",
			"key":        "Email",
			"sortable":   true,
			"searchable": true,
		},
	}

	var totalUsers int64
	models.DB.Model(&models.User{}).Where("Username != ?", "admin").Count(&totalUsers)

	responseData := map[string]interface{}{
		"columnsConfig": columnsConfig,
		"items":         users,
		"pagination": map[string]int{
			"currentPage": body.CurrentPage,
			"pageSize":    body.PageSize,
			"totalItems":  int(totalUsers),
		},
	}
	SendResponse(c, http.StatusOK, 200, responseData)
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

	if userReq.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
		if err != nil {
			SendResponse(c, http.StatusOK, apierrors.InternalServerError, nil)
			return
		}
		userReq.User.Password = string(hashedPassword)
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

func GetCurrentUser(c *gin.Context) {
	var user models.User

	// 从Gin的上下文中获取当前用户的ID
	currentUserIdStr, ok := c.Get("user_id")
	if !ok {
		SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
		return
	}

	currentUserId, err := strconv.Atoi(currentUserIdStr.(string))
	if err != nil {
		SendResponse(c, http.StatusOK, apierrors.InternalServerError, nil)
		return
	}

	// 查询数据库获取用户信息
	if result := models.DB.Preload("Roles").First(&user, currentUserId); result.Error != nil {
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

func UpdateCurrentUser(c *gin.Context) {
	var userReq models.UserRequest

	currentUserIdStr, ok := c.Get("user_id")
	if !ok {
		SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
		return
	}

	currentUserId, err := strconv.Atoi(currentUserIdStr.(string))
	if err != nil {
		SendResponse(c, http.StatusOK, apierrors.InternalServerError, nil)
		return
	}

	// 查询数据库获取用户信息
	if result := models.DB.First(&userReq.User, currentUserId); result.Error != nil {
		SendResponse(c, http.StatusOK, apierrors.DataNotFound, nil)
		return
	}

	if err := c.BindJSON(&userReq); err != nil {
		SendResponse(c, http.StatusOK, apierrors.InvalidRequestData, nil)
		return
	}

	if userReq.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
		if err != nil {
			SendResponse(c, http.StatusOK, apierrors.InternalServerError, nil)
			return
		}
		userReq.User.Password = string(hashedPassword)
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
