package db

import (
	"context"
	"reflect"

	"github.com/cham/elk/internal/db/model"
	"gorm.io/gorm"
)

type CRUD struct {
	db *gorm.DB
}

func NewCRUD(db *gorm.DB) *CRUD {
	return &CRUD{db: db}
}

// User CRUD
func (c *CRUD) CreateUser(ctx context.Context, user *model.User) error {
	return c.db.WithContext(ctx).Create(user).Error
}

func (c *CRUD) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	err := c.db.WithContext(ctx).First(&user, id).Error
	return &user, err
}

func (c *CRUD) UpdateUser(ctx context.Context, user *model.User) error {
	return c.db.WithContext(ctx).Save(user).Error
}

func (c *CRUD) DeleteUser(ctx context.Context, id uint) error {
	return c.db.WithContext(ctx).Delete(&model.User{}, id).Error
}

// Schema CRUD
func (c *CRUD) CreateSchema(ctx context.Context, schema *model.Schema) error {
	return c.db.WithContext(ctx).Create(schema).Error
}

func (c *CRUD) GetSchemaByID(ctx context.Context, id uint) (*model.Schema, error) {
	var schema model.Schema
	err := c.db.WithContext(ctx).First(&schema, id).Error
	return &schema, err
}

func (c *CRUD) UpdateSchema(ctx context.Context, schema *model.Schema) error {
	return c.db.WithContext(ctx).Save(schema).Error
}

func (c *CRUD) DeleteSchema(ctx context.Context, id uint) error {
	return c.db.WithContext(ctx).Delete(&model.Schema{}, id).Error
}

// Code CRUD
func (c *CRUD) CreateCode(ctx context.Context, code *model.Code) error {
	return c.db.WithContext(ctx).Create(code).Error
}

func (c *CRUD) GetCodeByID(ctx context.Context, id uint) (*model.Code, error) {
	var code model.Code
	err := c.db.WithContext(ctx).First(&code, id).Error
	return &code, err
}

func (c *CRUD) UpdateCode(ctx context.Context, code *model.Code) error {
	return c.db.WithContext(ctx).Save(code).Error
}

func (c *CRUD) DeleteCode(ctx context.Context, id uint) error {
	return c.db.WithContext(ctx).Delete(&model.Code{}, id).Error
}

// User列表查询
func (c *CRUD) ListUsers(ctx context.Context, params QueryParams) ([]model.User, int64, error) {
	var users []model.User
	var total int64
	query := c.db.WithContext(ctx).Model(&model.User{})

	// 应用查询条件
	query = applyQueryParams(query, params)

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页
	if params.Page > 0 && params.PageSize > 0 {
		query = query.Offset((params.Page - 1) * params.PageSize).Limit(params.PageSize)
	}

	// 执行查询
	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Schema列表查询
func (c *CRUD) ListSchemas(ctx context.Context, params QueryParams) ([]model.Schema, int64, error) {
	var schemas []model.Schema
	var total int64
	query := c.db.WithContext(ctx).Model(&model.Schema{})

	query = applyQueryParams(query, params)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if params.Page > 0 && params.PageSize > 0 {
		query = query.Offset((params.Page - 1) * params.PageSize).Limit(params.PageSize)
	}

	if err := query.Find(&schemas).Error; err != nil {
		return nil, 0, err
	}

	return schemas, total, nil
}

// Code列表查询
func (c *CRUD) ListCodes(ctx context.Context, params QueryParams) ([]model.Code, int64, error) {
	var codes []model.Code
	var total int64
	query := c.db.WithContext(ctx).Model(&model.Code{})

	query = applyQueryParams(query, params)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if params.Page > 0 && params.PageSize > 0 {
		query = query.Offset((params.Page - 1) * params.PageSize).Limit(params.PageSize)
	}

	if err := query.Find(&codes).Error; err != nil {
		return nil, 0, err
	}

	return codes, total, nil
}

// 应用查询参数
func applyQueryParams(query *gorm.DB, params QueryParams) *gorm.DB {
	if len(params.Fields) > 0 {
		query = query.Select(params.Fields)
	}

	for key, value := range params.Filters {
		switch v := value.(type) {
		case string:
			query = query.Where(key+" LIKE ?", "%"+v+"%")
		case map[string]interface{}:
			if v["lt"] != nil {
				query = query.Where(key+" < ?", v["lt"])
			}
			if v["gt"] != nil {
				query = query.Where(key+" > ?", v["gt"])
			}
			if v["gte"] != nil {
				query = query.Where(key+" >= ?", v["gte"])
			}
			if v["lte"] != nil {
				query = query.Where(key+" <= ?", v["lte"])
			}
		default:
			query = query.Where(key+" = ?", v)
		}
	}

	return query
}

// 辅助函数：检查值是否为零值
func isZero(v interface{}) bool {
	return v == nil || reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}

func (c *CRUD) GetUserByUsername(ctx context.Context, name string) (*model.User, error) {
	var user model.User
	err := c.db.WithContext(ctx).Where("username = ?", name).First(&user).Error
	return &user, err
}
