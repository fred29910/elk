package user

import (
	"github.com/cham/elk/internal/model/user"
	"github.com/cham/elk/internal/ov"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// 初始化数据库连接
	// db = ...
}

func List(c *gin.Context, query ov.UserQuery) ([]user.User, int64, error) {
	var users []user.User
	var total int64

	tx := db.Model(&user.User{})

	// 应用查询条件
	if query.Username != "" {
		tx = tx.Where("username LIKE ?", "%"+query.Username+"%")
	}
	if query.Email != "" {
		tx = tx.Where("email LIKE ?", "%"+query.Email+"%")
	}

	// 计算总数
	tx.Count(&total)

	// 排序
	if query.SortBy != "" {
		order := "ASC"
		if query.SortDesc {
			order = "DESC"
		}
		tx = tx.Order(query.SortBy + " " + order)
	}

	// 分页
	offset := (query.Page - 1) * query.PageSize
	tx = tx.Offset(offset).Limit(query.PageSize)

	result := tx.Find(&users)
	return users, total, result.Error
}
