/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-21 10:50:43
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-21 10:55:29
 * @FilePath: /go-mall/serializer/money.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import (
	"go-mall/model"
	"go-mall/pkg/util"
)

type Money struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(item *model.User, key string) Money {
	util.Encrypt.SetKey(key)
	return Money{
		UserId:    item.ID,
		UserName:  item.UserName,
		UserMoney: util.Encrypt.AesDecoding(item.Money),
	}
}
