package handlers

import (
	"gateway/apierrors"
	"gateway/dao"
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ruleDAO = dao.CasbinRuleDAO{DB: models.DB}

func GetPolicies(c *gin.Context) {
	rules, err := ruleDAO.GetPolicies()
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{
			Err:  apierrors.NewCustomError(apierrors.DatabaseError), // 这里可以选择适当的错误代码，这只是一个例子
			Type: gin.ErrorTypePublic,
		})
		return
	}
	SendResponse(c, http.StatusOK, apierrors.Success, rules)
}

func AddPolicy(c *gin.Context) {
	var rule models.CasbinRule
	if err := c.ShouldBindJSON(&rule); err != nil {
		c.Errors = append(c.Errors, &gin.Error{
			Err:  apierrors.NewCustomError(apierrors.InvalidRequestData),
			Type: gin.ErrorTypeBind,
		})
		return
	}

	if err := ruleDAO.AddPolicy(rule); err != nil {
		c.Errors = append(c.Errors, &gin.Error{
			Err:  apierrors.NewCustomError(apierrors.DatabaseError), // 同样，这里可以选择适当的错误代码
			Type: gin.ErrorTypePublic,
		})
		return
	}

	SendResponse(c, http.StatusOK, apierrors.Success, rule)
}
