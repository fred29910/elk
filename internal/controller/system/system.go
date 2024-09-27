package system

import (
	"github.com/cham/elk/internal/controller"
	"github.com/cham/elk/internal/dao/system"
	modelSts "github.com/cham/elk/internal/model/system"
	"github.com/cham/elk/pkg/ov"
	"github.com/gin-gonic/gin"
)

var _ = modelSts.System{}

// List
// @Summary 获取系统config列表
// @Description 获取系统config列表
// @Tags 系统管理
// @Accept json
// @Produce json
// @Param query query ov.SystemQuery true "查询参数"
// @Success 200 {object} ov.PageInfo{data=[]modelSts.System} "系统config列表"
// @Failure 400 {object} ov.Error
// @Router /api/systems [get]
func List(c *gin.Context) {
	var query ov.SystemQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	systems, total, err := system.List(c, query)
	if err != nil {
		c.Error(err)
		return
	}
	controller.PageResult(c, systems, total)
}

// Create
// @Summary 创建系统config
// @Description 创建系统config
// @Tags 系统管理
// @Accept json
// @Produce json
// @Param data body ov.CreateSystem true "系统config信息"
// @Success 200 {object} ov.Result{data=modelSts.System} "创建系统config"
// @Failure 400 {object} ov.Error
// @Router /api/systems [post]
func Create(c *gin.Context) {
	var data ov.CreateSystem
	if err := c.ShouldBindJSON(&data); err != nil {
		c.Error(err)
		return
	}

	system, err := system.Create(c, data)
	if err != nil {
		c.Error(err)
		return
	}
	controller.Success(c, system)
}
