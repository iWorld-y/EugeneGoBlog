package utils

import (
	gojwt "github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

var jwtKey []byte

type Claims struct {
	Uid int
	gojwt.StandardClaims
}

func init() {
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
}

// Award 生成 Token
func Award(uid *int) (string, error) {
	// 默认过期时间: 30 秒
	expireTime := time.Now().Add(time.Second * 30)
	claims := &Claims{
		Uid: *uid,
		StandardClaims: gojwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// 生成 Token
	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)
	if tokenStr, err := token.SignedString(jwtKey); err == nil {
		// 无异常
		return tokenStr, nil
	} else {
		log.Println(err)
		return "", err
	}
}

// ParseToken 解析 Token
func ParseToken(tokenStr string) (*gojwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := gojwt.ParseWithClaims(tokenStr, claims, func(t *gojwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, nil, err
	}
	return token, claims, nil
}
