package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	JwtKey                 = []byte("admin")
	TokenExpireTime        = time.Now().Add(time.Second * 3600 * 24).Unix()
	RefreshTokenExpireTime = time.Now().Add(time.Second * 3600 * 24).Unix()
)

type UserClaims struct {
	ID      uint
	Name    string
	IsAdmin bool
	jwt.RegisteredClaims
}

func GenerateToken(id uint, name string, expireTime int64) (string, error) {
	uc := UserClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expireTime))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		panic(err)
	}
	return tokenString, nil
}
