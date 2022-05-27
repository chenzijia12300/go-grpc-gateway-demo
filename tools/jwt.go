package tools

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	request "grpc-demo/model/request"
	"time"
)

var (
	secretKey        = []byte("secret")
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func GenerateToken(username string, id uint64) (string, error) {
	nowSec := time.Now().Unix()
	expired := time.Hour * 24 * 7
	claims := request.BaseClaims{
		ID:       id,
		Username: username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, request.CustomClaims{
		BaseClaims: claims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: nowSec + int64(expired.Seconds()), // 过期时间 7天  配置文件
			Issuer:    "grpc.demo",
		},
	})
	return token.SignedString(secretKey)
}

func ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return secretKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token == nil {
		return nil, TokenInvalid
	}
	return token.Claims.(*request.CustomClaims), nil
}
