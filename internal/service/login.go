package service

import (
	"context"
	"errors"
	"time"

	"github.com/cham/elk/internal/db"
	"github.com/cham/elk/internal/db/model"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	crud *db.CRUD
}

func NewLoginService(crud *db.CRUD) *LoginService {
	return &LoginService{crud: crud}
}

func (s *LoginService) Login(ctx context.Context, username, password string) (string, string, int, error) {
	user, err := s.crud.GetUserByUsername(ctx, username)
	if err != nil {
		return "", "", 0, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", 0, errors.New("密码不正确")
	}

	// 生成JWT令牌
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", "", 0, err
	}

	refreshToken := "refresh-token" // 这里应该生成一个真正的刷新令牌
	expiresIn := 86400              // 24小时

	return tokenString, refreshToken, expiresIn, nil
}

func (s *LoginService) Logout(ctx context.Context, token string) error {
	// 实现注销逻辑，可能需要将令牌加入黑名单
	return nil
}

func (s *LoginService) CurrentUser(ctx context.Context, token string) (*model.User, error) {
	// 解析JWT令牌
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		return nil, err
	}

	userID := uint(claims["user_id"].(float64))
	return s.crud.GetUserByID(ctx, userID)
}

func (s *LoginService) ListUsers(ctx context.Context, fields []string, filters map[string]interface{}, page, pageSize int) ([]model.User, int64, error) {

	queryParams := db.QueryParams{
		Fields:   fields,
		Filters:  filters,
		Page:     page,
		PageSize: pageSize,
	}

	return s.crud.ListUsers(ctx, queryParams)
}
