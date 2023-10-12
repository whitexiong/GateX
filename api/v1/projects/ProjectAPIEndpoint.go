package projects

import (
	"gateway/api/v1/setting"
	"gateway/dao"
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var endpointDAO = dao.NewProjectAPIEndpointDAO(models.DB)

func ListProjectAPIEndpoints(c *gin.Context) {
	projectId, err := strconv.ParseUint(c.Param("projectId"), 10, 32)
	if err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid project ID.")
		return
	}

	endpoints, err := endpointDAO.ListEndpointsByProject(uint(projectId))
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to retrieve API endpoints for the project.")
		return
	}
	setting.SendResponse(c, http.StatusOK, 200, endpoints)
}

func CreateProjectAPIEndpoint(c *gin.Context) {
	projectId, err := strconv.ParseUint(c.Param("projectId"), 10, 32)
	if err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid project ID.")
		return
	}

	var endpoint models.ProjectAPIEndpoint
	if err := c.ShouldBindJSON(&endpoint); err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid endpoint data.")
		return
	}
	endpoint.ProjectID = uint(projectId)

	err = endpointDAO.CreateEndpoint(&endpoint)
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to create API endpoint.")
		return
	}
	setting.SendResponse(c, http.StatusOK, 200, "API endpoint created successfully.")
}

func GetProjectAPIEndpointDetails(c *gin.Context) {
	endpointId, err := strconv.ParseUint(c.Param("endpointId"), 10, 32)
	if err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid endpoint ID.")
		return
	}

	endpoint, err := endpointDAO.GetEndpointByID(uint(endpointId))
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to retrieve API endpoint details.")
		return
	}
	setting.SendResponse(c, http.StatusOK, 200, endpoint)
}

func UpdateProjectAPIEndpoint(c *gin.Context) {
	endpointId, err := strconv.ParseUint(c.Param("endpointId"), 10, 32)
	if err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid endpoint ID.")
		return
	}

	var endpoint models.ProjectAPIEndpoint
	if err := c.ShouldBindJSON(&endpoint); err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid endpoint data.")
		return
	}

	endpoint.ID = uint(endpointId)
	err = endpointDAO.UpdateEndpoint(&endpoint)
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to update API endpoint.")
		return
	}
	setting.SendResponse(c, http.StatusOK, 200, "API endpoint updated successfully.")
}

func DeleteProjectAPIEndpoint(c *gin.Context) {
	endpointId, err := strconv.ParseUint(c.Param("endpointId"), 10, 32)
	if err != nil {
		setting.SendResponse(c, http.StatusBadRequest, -1, "Invalid endpoint ID.")
		return
	}

	err = endpointDAO.DeleteEndpoint(uint(endpointId))
	if err != nil {
		setting.SendResponse(c, http.StatusInternalServerError, -1, "Failed to delete API endpoint.")
		return
	}
	setting.SendResponse(c, http.StatusOK, 200, "API endpoint deleted successfully.")
}
