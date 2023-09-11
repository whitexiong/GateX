package apierrors

import "github.com/gin-gonic/gin"

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

const (
	Success              = 200
	InvalidRequestData   = 400
	Unauthorized         = 401
	PermissionDenied     = 403
	DataNotFound         = 404
	MethodNotAllowed     = 405
	Conflict             = 409
	InternalServerError  = 500
	ServiceUnavailable   = 503
	DatabaseError        = 600
	ValidationError      = 601
	AuthenticationFailed = 602
	RateLimitExceeded    = 603
)

var errorDescriptions = map[int]string{
	Success:              "success",
	InvalidRequestData:   "无效的请求数据",
	Unauthorized:         "未经授权",
	PermissionDenied:     "权限被拒绝",
	DataNotFound:         "数据未找到",
	MethodNotAllowed:     "方法不允许",
	Conflict:             "数据冲突",
	InternalServerError:  "服务器内部错误",
	ServiceUnavailable:   "服务不可用",
	DatabaseError:        "数据库错误",
	ValidationError:      "验证错误",
	AuthenticationFailed: "认证失败",
	RateLimitExceeded:    "超出频率限制",
}

func NewCustomError(code int) *CustomError {
	return &CustomError{
		Code:    code,
		Message: GetDescription(code),
	}
}

func HandleGinError(c *gin.Context, errorCode int) {
	errMsg := NewCustomError(errorCode)
	c.Errors = append(c.Errors, &gin.Error{
		Err:  errMsg,
		Type: gin.ErrorTypePublic,
	})
	c.Abort()
}

func GetDescription(code int) string {
	if msg, ok := errorDescriptions[code]; ok {
		return msg
	}
	return "success"
}
