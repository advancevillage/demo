//author: richard
package util

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signature = "d41d8cd98f00b204e9800998ecf8427e"
	expire = 7
	iss = "demo"
	ErrorFetchValueFromTokenFormat = "fetch %s from %s failed"
)
//@brief: 创建Token
func AuthCreateTokenService(id string) (token string, err error) {
	j := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims["iat"] = time.Now().Unix() //发行开始时间
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24 * expire)).Unix()
	claims["iss"] = iss
	claims["jti"] = RandomString(12)
	claims["nbf"] = time.Date(2019,6,1,0,0,0, 0, time.UTC).Unix()
	j.Claims = claims
	token, err = j.SignedString([]byte(signature))
	return
}
//@brief: 解析Token
func AuthParseTokenService(token string) (b bool) {
	j, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(signature), nil
	})
	if err != nil || !j.Valid {
		b = false
	} else {
		b = true
	}
	return
}
//@brief: 获取值从Token
func AuthFetchValueFromToken(token string, key string) (value string, err error) {
	j, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(signature), nil
	})
	if err != nil || !j.Valid {
		return "", err
	}
	claims := j.Claims.(jwt.MapClaims)
	if value, ok := claims[key]; ok {
		return fmt.Sprintf("%v", value), nil
	} else {
		return "", errors.New(fmt.Sprintf(ErrorFetchValueFromTokenFormat,key, token))
	}
}



