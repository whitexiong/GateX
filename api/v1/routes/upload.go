package routes

import (
	v1 "gateway/api/v1/setting"
	"github.com/gin-gonic/gin"
)

func SetupUploads(r *gin.Engine) {
	r.POST("/upload", v1.UploadFile)
}
