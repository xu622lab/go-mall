/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-17 10:19:13
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-18 11:14:51
 * @FilePath: /go-mall/serializer/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import (
	"go-mall/model"
)

// VO
type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nickname"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

// BuildUser 序列化用户
func BuildUser(user *model.User) *User {
	u := &User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}

	return u
}
