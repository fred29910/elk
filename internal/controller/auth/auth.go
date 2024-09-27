package auth

import (
	"net/http"

	"github.com/cham/elk/internal/dao/user"
	"github.com/cham/elk/pkg/errors"
	"github.com/cham/elk/pkg/ov"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary 登陆
// @Description 登陆
// @Tags auth
// @Accept json
// @Produce json
// @Param body body ov.LoginRequest true "登陆参数"
// @Success 200 {object} ov.Token
// @Failure 400 {object} ov.Error
// @Router /api/auth/login [post]
func Login(c *gin.Context) {
	var data ov.LoginRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(errors.VALID_ERROR)
		return
	}

	// 查询用户信息
	user, err := user.Get(c, map[string]any{
		"username": data.Username,
	})
	if err != nil {
		c.Error(errors.USER_NOT_FOUND)
		return
	}

	if user.Password != data.Password {
		c.Error(errors.LOGIN_ERROR)
		return
	}

	c.JSON(http.StatusOK, &ov.Token{
		Token:        uuid.New().String(),
		RefreshToken: uuid.New().String(),
		ExpiresIn:    3600,
	})
}

// @Summary 刷新token
// @Description 刷新token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body ov.RefreshTokenRequest true "刷新token参数"
// @Success 200 {object} ov.Token
// @Failure 400 {object} ov.Error
// @Router /api/auth/refresh_token [post]
func RefreshToken(c *gin.Context) {
	var data ov.RefreshTokenRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(errors.VALID_ERROR)
		return
	}

}

// @Summary 登出
// @Description 登出
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {string} string "登出成功"
// @Failure 400 {object} ov.Error
// @Router /api/auth/logout [post]
func Logout(c *gin.Context) {

}
