/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-17 09:55:16
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-17 10:11:50
 * @FilePath: /go-mall/pkg/util/jwt.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// token签名用
var jwtSecret = []byte("yijiansanlian")

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

// 签发token
func GenerateToken(id uint, userName string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		ID:        id,
		UserName:  userName,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "XZY-Mall",
		},
	}

	// 加密一下
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签发
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
