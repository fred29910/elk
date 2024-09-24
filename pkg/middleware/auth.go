package middleware

import (
	"net/http"

	"github.com/cham/elk/pkg/core"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里应该实现您的身份验证逻辑
		// 例如,检查 JWT token, session 等

		// 示例: 从请求头中获取 token
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供认证令牌"})
			c.Abort()
			return
		}

		// TODO: 验证 token 并获取用户信息
		// 这里应该包含实际的 token 验证逻辑

		// 示例: 假设我们已经验证了 token 并获取了用户信息
		userInfo := &core.ContextUserInfo{
			ID:       "user123",
			Username: "exampleUser",
			Email:    "user@example.com",
		}

		// 将用户信息存储到上下文中
		core.StoreContextUserInfo(c, userInfo)

		c.Next()
	}
}
