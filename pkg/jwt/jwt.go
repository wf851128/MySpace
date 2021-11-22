package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var TokenSecretKey = "I love more than I can say"

// AccessTokenExpireDuration aToken失效时间
const AccessTokenExpireDuration = time.Hour * 6

// RefreshTokenExpireDuration rToken失效时间
const RefreshTokenExpireDuration = time.Hour * 24 * 30

//IssuerKey 签发人
const IssuerKey = "MySpace"

var MySecret = []byte(TokenSecretKey)

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

// GenToken 生成 token
func GenToken(userId int64, username string) (aToken, rToken string, err error) {
	//定义access Token，带有用户信息的结构体
	c := MyClaims{
		UserID:   userId,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpireDuration)),
			Issuer:    IssuerKey,
		},
	}
	//生成 access Token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(MySecret)
	//生成 refresh Token，没有自定义数据信息的
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenExpireDuration)),
		Issuer:    IssuerKey,
	}).SignedString(MySecret)

	return aToken, rToken, err
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

func RefreshToken(aToken, rToken string) (string, error) {
	//检查 refresh token 是否有效
	_, err := jwt.Parse(rToken, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return "", err
	}
	//	从旧的 access token 中解析出自定义数据,传给claims
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	//验证 aToken 是否是因为超期产生错误
	v, _ := err.(*jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		//如果是因为超期而产生的错误，则新生产的 aToken
		newAToken, _, _ := GenToken(claims.UserID, claims.UserName)
		return newAToken, nil
	}
	return "", err
}
