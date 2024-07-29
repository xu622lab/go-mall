/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-29 10:26:49
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-29 11:00:16
 * @FilePath: /go-mall/serializer/address.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import "go-mall/model"

type Address struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	CreateAt int64  `json:"create_at"`
}

func BuildAddress(item *model.Address) Address {
	return Address{
		Id:       item.ID,
		UserId:   item.UserID,
		Name:     item.Name,
		Phone:    item.Phone,
		Address:  item.Address,
		CreateAt: item.CreatedAt.Unix(),
	}
}

func BuildAddreses(items []*model.Address) (addresses []Address) {
	for _, item := range items {
		address := BuildAddress(item)
		addresses = append(addresses, address)
	}

	return addresses
}
