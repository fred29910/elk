package middleware

import (
	"net/http"

	"github.com/cham/elk/internal/dao/user"
	"github.com/cham/elk/internal/service/auth"
	"github.com/cham/elk/pkg/core"
	"github.com/cham/elk/pkg/errors"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 验证用户身份
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里应该实现您的身份验证逻辑
		// 路径白名单
		path := c.FullPath()
		if path == "/api/auth/login" || path == "/api/auth/register" {
			c.Next()
			return
		}
		// 例如,检查 JWT token, session 等

		// 示例: 从请求头中获取 token
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供认证令牌"})
			c.Abort()
			return
		}

		// TODO: 验证 token 并获取用户信息
		uid, err := auth.VerifyToken(token)
		if err != nil {
			c.Error(errors.UNAUTHORIZED)
			c.Abort()
			return
		}
		user, err := user.Get(c, map[string]any{
			"id": uid,
		})
		if err != nil {
			c.Error(errors.USER_NOT_FOUND)
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		core.StoreContextUserInfo(c, &core.ContextUserInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})

		c.Next()
	}
}
