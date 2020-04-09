// @Author:WY
// @Time:2020/4/7 16:13
package utils

import (
	"gin_vue/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtkey = []byte("a_secret_crect")

type Claims struct {
	UserId int
	jwt.StandardClaims
}

func ReleaseToken(user models.User) (string, error) {
	//设置token有效时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //到期时间
			IssuedAt:  time.Now().Unix(),     //发放时间
			Issuer:    "www.wenicer.com",     //发放机构
			Subject:   "user token",          //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStrng, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenStrng, nil
}

func ParseToken(tokenString string)(*jwt.Token,*Claims,error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	return token, claims, err
}
