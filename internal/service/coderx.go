package service

import (
	"context"

	"github.com/cham/elk/internal/db"
	"github.com/cham/elk/internal/db/model"
)

type CoderXService struct {
	crud *db.CRUD
}

func NewCoderXService(crud *db.CRUD) *CoderXService {
	return &CoderXService{crud: crud}
}

func (s *CoderXService) CreateUser(ctx context.Context, user *model.User) error {
	return s.crud.CreateUser(ctx, user)
}

func (s *CoderXService) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	return s.crud.GetUserByID(ctx, id)
}

func (s *CoderXService) UpdateUser(ctx context.Context, user *model.User) error {
	return s.crud.UpdateUser(ctx, user)
}

func (s *CoderXService) DeleteUser(ctx context.Context, id uint) error {
	return s.crud.DeleteUser(ctx, id)
}

func (s *CoderXService) ListUsers(ctx context.Context, params db.QueryParams) ([]model.User, int64, error) {
	return s.crud.ListUsers(ctx, params)
}

func (s *CoderXService) CreateSchema(ctx context.Context, schema *model.Schema) error {
	return s.crud.CreateSchema(ctx, schema)
}

func (s *CoderXService) GetSchemaByID(ctx context.Context, id uint) (*model.Schema, error) {
	return s.crud.GetSchemaByID(ctx, id)
}

func (s *CoderXService) UpdateSchema(ctx context.Context, schema *model.Schema) error {
	return s.crud.UpdateSchema(ctx, schema)
}

func (s *CoderXService) DeleteSchema(ctx context.Context, id uint) error {
	return s.crud.DeleteSchema(ctx, id)
}

func (s *CoderXService) ListSchemas(ctx context.Context, params db.QueryParams) ([]model.Schema, int64, error) {
	return s.crud.ListSchemas(ctx, params)
}

func (s *CoderXService) CreateCode(ctx context.Context, code *model.Code) error {
	return s.crud.CreateCode(ctx, code)
}

func (s *CoderXService) GetCodeByID(ctx context.Context, id uint) (*model.Code, error) {
	return s.crud.GetCodeByID(ctx, id)
}

func (s *CoderXService) UpdateCode(ctx context.Context, code *model.Code) error {
	return s.crud.UpdateCode(ctx, code)
}

func (s *CoderXService) DeleteCode(ctx context.Context, id uint) error {
	return s.crud.DeleteCode(ctx, id)
}

func (s *CoderXService) ListCodes(ctx context.Context, params db.QueryParams) ([]model.Code, int64, error) {
	return s.crud.ListCodes(ctx, params)
}

func (s *CoderXService) GetUserByUsername(ctx context.Context, name string) (*model.User, error) {
	return s.crud.GetUserByUsername(ctx, name)
}
