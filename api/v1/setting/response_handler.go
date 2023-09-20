package setting

import (
	"gateway/apierrors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func SendResponse(c *gin.Context, httpCode int, code int, data interface{}) {
	msg := apierrors.GetDescription(code)
	c.JSON(httpCode, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
