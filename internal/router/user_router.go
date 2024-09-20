package router

import (
	"github.com/cham/elk/internal/controller/user"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/list", user.List)
		userGroup.POST("/create", user.Create)
		// 添加更多路由...
	}
}
