package handlers

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// 统一相应
func SendResponse(c *gin.Context, httpCode int, code int, data interface{}, msg string) {
	c.JSON(httpCode, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
