package system

import (
	"github.com/cham/elk/internal/dao"
	"github.com/cham/elk/internal/model/system"
	"github.com/cham/elk/pkg/ov"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context, query ov.SystemQuery) ([]system.System, int64, error) {
	var systems []system.System
	var total int64

	tx := dao.DB(c).Model(&system.System{})

	// 应用查询条件
	if query.Type != nil {
		tx = tx.Where("type = ?", *query.Type)
	}
	if query.Status != nil {
		tx = tx.Where("status = ?", *query.Status)
	}

	// count
	tx.Count(&total)

	// sort
	if query.SortBy != nil {
		order := "ASC"
		if query.SortDesc != nil && *query.SortDesc {
			order = "DESC"
		}
		tx = tx.Order(*query.SortBy + " " + order)
	}

	// page
	tx = dao.Pageble(tx, query.Page, query.PageSize)

	result := tx.Find(&systems)
	return systems, total, result.Error
}

// create
func Create(c *gin.Context, data ov.CreateSystem) (*system.System, error) {
	system := system.System{
		Config: data.Config,
		Type:   data.Type,
		Status: data.Status,
		Remark: data.Remark,
	}
	err := dao.DB(c).Create(&system).Error
	return &system, err
}
