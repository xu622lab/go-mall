/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-16 11:27:42
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-16 11:28:11
 * @FilePath: /go-mall/model/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import (
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `gorm:"size:1000"`
	Money          string
}
