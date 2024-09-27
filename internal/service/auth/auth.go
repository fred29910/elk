package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/cham/elk/internal/config"
	"github.com/cham/elk/internal/dao/token"
	"github.com/golang-jwt/jwt/v5"
)

// create token by uid
func CreateToken(uid uint) (string, string, error) {
	jwtConfig := config.GetJwtConfig()
	tokenSp := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
	})
	tokenString, err := tokenSp.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return "", "", err

	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		return "", "", err
	}

	// set token to cache, default expiration is 1 hour
	token.Set(tokenString, uid, time.Duration(jwtConfig.Expiration)*time.Second)
	token.Set(refreshTokenString, uid, time.Duration(jwtConfig.RefreshExpiration)*time.Second)

	token.Set(fmt.Sprintf("user:%d:token", uid), tokenString, time.Duration(jwtConfig.Expiration)*time.Second)
	token.Set(fmt.Sprintf("user:%d:refresh_token", uid), refreshTokenString, time.Duration(jwtConfig.RefreshExpiration)*time.Second)
	return tokenString, refreshTokenString, nil
}

// verify token
func VerifyToken(tokenString string) (uint, error) {

	uid, err := token.Get(tokenString)
	if err {
		return uid.(uint), nil
	}
	return 0, errors.New("token not found")
}

func Logout(tokenString string) error {

	uid, err := VerifyToken(tokenString)
	if err != nil {
		return err
	}
	// get refresh token
	refreshToken, ok := token.Get(fmt.Sprintf("user:%d:refresh_token", uid))
	if !ok {
		return errors.New("refresh token not found")
	}

	token.Delete(tokenString)
	token.Delete(refreshToken.(string))
	token.Delete(fmt.Sprintf("user:%d:token", uid))
	token.Delete(fmt.Sprintf("user:%d:refresh_token", uid))
	return nil
}
