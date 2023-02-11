package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"go_learn/app/model"
	"time"
)

type AdminClaims struct {
	UserId   int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte(viper.GetString("server.jwt_secret"))
var exp = time.Hour * 4

func GenerateToken(u model.Admin) (string, error) {
	expTime := time.Now().Add(exp).Unix()
	claims := &AdminClaims{
		u.Id,
		u.Username,
		jwt.StandardClaims{
			ExpiresAt: expTime,
			Issuer:    "gin-jwt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *AdminClaims, error) {
	claims := &AdminClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}