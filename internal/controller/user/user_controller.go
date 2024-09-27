package user

import (
	"github.com/cham/elk/internal/service/user"
	"github.com/cham/elk/pkg/ov"

	"github.com/gin-gonic/gin"
)

// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param query query ov.UserQuery true "查询参数"
// @Success 200 {object} ov.User
// @Failure 400 {object} ov.Error
// @Router /api/users [get]
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

// @Summary 创建用户
// @Description 创建用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body ov.User true "用户信息"
// @Success 200 {object} ov.User
// @Failure 400 {object} ov.Error
// @Router /api/users [post]
func Create(c *gin.Context) {
	var data ov.User
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := user.Create(c, data)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, createdUser)
}
