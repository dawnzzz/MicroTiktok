package auth

import (
	"errors"
	"fmt"
	"github.com/dawnzzz/MicroTiktok/internal/service/authentication/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaims struct {
	userID int64
	jwt.StandardClaims
}

// GetToken 获取token
func GetToken(userID int64) (string, error) {
	cla := UserClaims{
		userID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.AuthenticationConfigObj.TokenExpireDuration).Unix(), // 过期时间
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	fmt.Println("Token = ", token)
	return token.SignedString(config.AuthenticationConfigObj.Secret) // 进行签名生成对应的token
}

// ParseToken 解析token
func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.AuthenticationConfigObj.Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
