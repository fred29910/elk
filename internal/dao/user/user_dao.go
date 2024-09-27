package user

import (
	"github.com/cham/elk/internal/dao"
	"github.com/cham/elk/internal/model/user"
	"github.com/cham/elk/internal/ov"
	pg "github.com/cham/elk/pkg/ov"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context, query ov.UserQuery) ([]user.User, int64, error) {
	var users []user.User
	var total int64

	tx := dao.DB(c).Model(&user.User{})

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

func Create(c *gin.Context, data pg.User) (*user.User, error) {
	user := user.User{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Status:   data.Status,
	}
	err := dao.DB(c).Create(&user).Error
	return &user, err
}

func Update(c *gin.Context, data pg.UpdateUserOV) (*user.User, error) {
	updataMap := make(map[string]interface{})
	if data.Username != nil {
		updataMap["username"] = data.Username
	}
	if data.Email != nil {
		updataMap["email"] = data.Email
	}
	if data.Password != nil {
		updataMap["password"] = data.Password
	}
	if data.Status != nil {
		updataMap["status"] = data.Status
	}
	err := dao.DB(c).Model(&user.User{}).Where("id = ?", data.ID).Updates(updataMap).Error
	if err != nil {
		return nil, err
	}

	var updatedUser user.User
	err = dao.DB(c).Model(&user.User{}).Where("id = ?", data.ID).First(&updatedUser).Error
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

func Get(c *gin.Context, qs map[string]any) (*user.User, error) {
	var user user.User

	err := dao.DB(c).Where(qs).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
