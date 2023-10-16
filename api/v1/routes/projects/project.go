package projects

import (
	"gateway/api/v1/projects"
	"github.com/gin-gonic/gin"
)

func SetupProjects(r *gin.Engine) {
	group := r.Group("/projects")
	group.POST("/list", projects.GetProjectList)
	group.POST("/add", projects.CreateProject)
	group.GET("/details/:id", projects.GetProjectDetails)
	group.POST("/update/:id", projects.UpdateProject)
	group.GET("/delete/:id", projects.DeleteProject)
}
