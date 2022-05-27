package request

import "github.com/golang-jwt/jwt/v4"

type CustomClaims struct {
	BaseClaims
	jwt.StandardClaims
}

type BaseClaims struct {
	ID       uint64
	Username string
}
