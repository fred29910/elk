package controller

import (
	"github.com/cham/elk/pkg/ov"
	"github.com/gin-gonic/gin"
)

// Success
func Success(c *gin.Context, data any) {
	c.JSON(200, ov.Result{
		Data:  data,
		Error: "",
	})
}

func Error(c *gin.Context, code int, err error) {
	c.JSON(200, ov.Result{
		Data:  nil,
		Code:  code,
		Error: err.Error(),
	})
}

func ErrorWithMsg(c *gin.Context, code int, msg string) {
	c.JSON(200, ov.Result{
		Data:  nil,
		Code:  code,
		Error: msg,
	})
}

func PageResult(c *gin.Context, data any, total int64) {
	c.JSON(200, ov.PageInfo{
		Data:  data,
		Total: total,
		Error: "",
	})
}
