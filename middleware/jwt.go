/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-18 10:27:43
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-18 10:48:24
 * @FilePath: /go-mall/middleware/jwt.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package middleware

import (
	"go-mall/pkg/e"
	"go-mall/pkg/util"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = 200
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}

		if code != e.Success {
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
