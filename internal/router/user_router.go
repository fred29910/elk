package router

import (
	"github.com/cham/elk/internal/controller/auth"
	"github.com/cham/elk/internal/controller/user"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine) {
	rg := r.Group("/api")
	// rg.Use(middleware.AuthMiddleware())
	userGroup := rg.Group("/users")
	{
		userGroup.GET("", user.List)
		userGroup.POST("", user.Create)
		// 添加更多路由...
	}
}

func InitAuthRoutes(r *gin.Engine) {
	rg := r.Group("/api")
	// rg.Use(middleware.AuthMiddleware())
	authGroup := rg.Group("/auth")
	{
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/register", auth.RefreshToken)
		// 添加更多路由...
	}
}
