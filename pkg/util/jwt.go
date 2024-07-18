/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-17 09:55:16
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-18 10:05:07
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

// 验证用户token
func ParseToken(token string) (*Claims, error) {
	// 解析token
	/*
		jwt.ParseWithClaims是一个方法，用于解析JWT（JSON Web Token）并将其映射到指定的声明结构体中。
		它接受三个参数：要解析的token字符串、一个声明结构体的实例指针以及一个回调函数。
		回调函数用于提供用于验证token的密钥。
		在解析过程中，它会验证token的签名和有效期，并尝试将token中的声明信息解析到提供的声明结构体实例中。
		如果解析成功且token有效，则返回一个*jwt.Token类型的对象，其中包含解析后的声明信息。
	*/
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	// 检查tokenClaims是否为非空，
	// 然后尝试将tokenClaims.Claims转换为*Claims类型的claims。如果转换成功且tokenClaims有效，则返回claims和nil，表示token验证成功。
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
