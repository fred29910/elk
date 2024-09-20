package user

import (
	"github.com/cham/elk/internal/ov"
	"github.com/cham/elk/internal/service/user"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var query ov.UserQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	users, total, err := user.List(c, query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data":  users,
		"total": total,
	})
}

func Create(c *gin.Context) {
	// 实现创建用户逻辑
	// ...
}
