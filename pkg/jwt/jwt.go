package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var TokenSecretKey = "I love more than I can say"

const TokenExpireDuration = time.Hour * 6

var MySecret = []byte(TokenSecretKey)

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

// GenToken 生成 token
func GenToken(userId int64, username string) (string, error) {
	c := MyClaims{
		UserID:   userId,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "MySpace",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*MyClaims, error) {
	//创建一个Claims的对象
	var mc = new(MyClaims)
	//将 tokenStr，使用 MySecret 进行解密，并将结果赋值给 mc
	token, err := jwt.ParseWithClaims(tokenStr, mc, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	//校验 token 是否正常
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
