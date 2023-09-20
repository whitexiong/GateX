package middleware

import (
	"errors"
	"gateway/api/v1/setting"
	"gateway/apierrors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func InitLoggingAndErrorHandlingMiddleware() gin.HandlerFunc {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next() // 处理请求

		latency := time.Since(startTime)
		status := c.Writer.Status()

		entry := logger.WithFields(logrus.Fields{
			"status":   status,
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"clientIP": c.ClientIP(),
			"latency":  latency,
		})

		if len(c.Errors) > 0 {
			// 如果有错误，发送统一的错误响应
			lastError := c.Errors.Last().Err // 获取最后一个错误

			var customErr *apierrors.CustomError
			if errors.As(lastError, &customErr) {
				setting.SendResponse(c, httpCodeBasedOnErrorCode(customErr.Code), customErr.Code, nil)
			}

			// 记录错误
			for _, e := range c.Errors {
				entry.WithField("error", e.Err).Error("Request failed")
			}
		} else {
			entry.Info("Request succeeded")
		}
	}
}

func httpCodeBasedOnErrorCode(code int) int {
	switch code {
	case apierrors.InvalidRequestData:
		return http.StatusBadRequest
	case apierrors.Unauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
