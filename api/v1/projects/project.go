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

var projectDAO = dao.NewProjectDAO(models.DB)

//func GetProjectList(c *gin.Context) {
//	projects, err := projectDAO.ListProjects()
//	if err != nil {
//		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to retrieve project list.")
//		return
//	}
//	setting.SendResponse(c, http.StatusOK, 200, projects)
//}

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
		// ... 添加其他需要的列配置
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

	err := projectDAO.CreateProject(&project)
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to create project.")
		return
	}
	setting.SendResponse(c, http.StatusOK, 200, "Project created successfully.")
}

func GetProjectDetails(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid project ID.")
		return
	}

	project, err := projectDAO.GetProjectByID(uint(id))
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
	err = projectDAO.UpdateProject(&project)
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

	err = projectDAO.DeleteProject(uint(id))
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to delete project.")
		return
	}
	setting.SendResponse(c, http.StatusOK, 200, "Project deleted successfully.")
}
