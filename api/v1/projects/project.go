package projects

import (
	"gateway/api/v1/setting"
	"gateway/dao"
	"gateway/models"
	"gateway/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProjectList(c *gin.Context) {
	type RequestBody struct {
		Name        string `json:"name"`
		PageSize    int    `json:"pageSize"`
		CurrentPage int    `json:"currentPage"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		setting.SendResponse(c, http.StatusBadRequest, 400, nil)
		return
	}

	paginator, searcher := pagination.NewPaginationService(body.CurrentPage, body.PageSize, body.Name,
		pagination.WithSearchField("Name"),
	)

	query := models.DB.Model(&models.Project{}).Preload("APIs")
	query = searcher.Search(query)
	query = paginator.Paginate(query)

	var projects []models.Project
	if result := query.Find(&projects); result.Error != nil {
		setting.SendResponse(c, http.StatusInternalServerError, 500, nil)
		return
	}

	columnsConfig := []map[string]interface{}{
		{
			"label":      "项目名称",
			"key":        "Name",
			"sortable":   true,
			"searchable": true,
		},
		{
			"label":      "描述",
			"key":        "Description",
			"sortable":   true,
			"searchable": true,
		},
	}

	var totalProjects int64
	models.DB.Model(&models.Project{}).Count(&totalProjects)

	responseData := map[string]interface{}{
		"columnsConfig": columnsConfig,
		"items":         projects,
		"pagination": map[string]int{
			"currentPage": body.CurrentPage,
			"pageSize":    body.PageSize,
			"totalItems":  int(totalProjects),
		},
	}
	setting.SendResponse(c, http.StatusOK, 200, responseData)
}

func CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid project data.")
		return
	}

	err := dao.DefaultProjectDAO.CreateProject(&project)
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to create project.")
		return
	}

	defaultSetting := &models.ProjectSetting{
		ProjectID:   project.ID,
		IP:          "127.0.0.1",
		Port:        80,
		Environment: "Development",
		Description: "Default setting for the project",
	}
	err = dao.DefaultProjectSettingDAO.CreateProjectSetting(defaultSetting)
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to create default project setting.")
		return
	}

	project.DefaultSettingID = defaultSetting.ID
	err = dao.DefaultProjectDAO.UpdateProject(&project)
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to update project with default setting ID.")
		return
	}

	setting.SendResponse(c, http.StatusOK, 200, "Project and default setting created successfully.")
}

func GetProjectDetails(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid project ID.")
		return
	}

	project, err := dao.DefaultProjectDAO.GetProjectByID(uint(id))
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to retrieve project details.")
		return
	}
	setting.SendResponse(c, http.StatusOK, 200, project)
}

func UpdateProject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid project ID.")
		return
	}

	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid project data.")
		return
	}

	project.ID = uint(id)
	err = dao.DefaultProjectDAO.UpdateProject(&project)
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to update project.")
		return
	}
	setting.SendResponse(c, http.StatusOK, 200, "Project updated successfully.")
}

func DeleteProject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid project ID.")
		return
	}

	err = dao.DefaultProjectDAO.DeleteProject(uint(id))
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to delete project.")
		return
	}
	setting.SendResponse(c, http.StatusOK, 200, "Project deleted successfully.")
}

func GetProjectSetting(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid project setting ID.")
		return
	}

	ProjectSetting, err := dao.DefaultProjectSettingDAO.GetProjectSettingByProjectID(uint(id))
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to retrieve project setting.")
		return
	}
	setting.SendResponse(c, http.StatusOK, 200, ProjectSetting)
}

func UpdateOrCreateProjectSetting(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid project setting ID.")
		return
	}

	var projectSetting models.ProjectSetting
	if err := c.ShouldBindJSON(&projectSetting); err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid project setting data.")
		return
	}

	projectSetting.ID = uint(id)
	err = dao.DefaultProjectSettingDAO.UpdateProjectSetting(&projectSetting)
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to update or create project setting.")
		return
	}
	setting.SendResponse(c, http.StatusOK, 200, "Project setting updated or created successfully.")
}
