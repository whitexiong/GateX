package setting

import (
	"gateway/apierrors"
	"gateway/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		SendResponse(c, http.StatusBadRequest, apierrors.InvalidRequestData, nil)
		return
	}

	savePath, err := upload.SaveUploadedFile(file, "./uploads")
	if err != nil {
		SendResponse(c, http.StatusInternalServerError, apierrors.UploadFileFail, nil)
		return
	}

	SendResponse(c, http.StatusOK, 200, gin.H{"url": savePath})
}
