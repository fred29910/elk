package router

import (
	"github.com/cham/elk/internal/controller/user"

	"github.com/cham/elk/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine) {
	rg := r.Group("/api")
	rg.Use(middleware.AuthMiddleware())
	userGroup := rg.Group("/users")
	{
		userGroup.GET("", user.List)
		userGroup.POST("", user.Create)
		// 添加更多路由...
	}
}
