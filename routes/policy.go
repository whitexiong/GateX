package routes

import (
	"gateway/handlers"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

func SetupPoliciesRoutes(r *gin.Engine) {
	menuGroup := r.Group("/policy")
	menuGroup.Use(middleware.InitJWTMiddleware())
	menuGroup.GET("list", handlers.GetPolicies)
	menuGroup.POST("add", handlers.AddPolicy)
	//menuGroup.DELETE("delete/:id", handlers.DeletePolicy)
	//menuGroup.PUT("update/:id", handlers.UpdatePolicy)
	//menuGroup.GET("detail/:id", handlers.DetailPolicy)

}
