/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-26 15:08:02
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-29 20:09:40
 * @FilePath: /go-mall/serializer/Order.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import (
	"context"
	"go-mall/conf"
	"go-mall/dao"
	"go-mall/model"
)

type Order struct {
	Id           uint    `json:"id"`
	OrderNum     uint64  `json:"order_num"`
	CreatedAt    int64   `json:"created_at"`
	UpdatedAt    int64   `json:"updated_at"`
	UserId       uint    `json:"user_id"`
	ProductId    uint    `json:"product_id"`
	BossId       uint    `json:"boss_id"`
	Num          int     `json:"num"`
	AddressName  string  `json:"address_name"`
	AddressPhone string  `json:"address_phone"`
	Address      string  `json:"address"`
	Type         uint    `json:"type"`
	ProductName  string  `json:"product_name"`
	ImgPath      string  `json:"img_path"`
	Money        float64 `json:"discount_price"`
}

func BuildOrder(order *model.Order, address *model.Address, product *model.Product) Order {
	return Order{
		Id:           order.ID,
		OrderNum:     order.OrderNum,
		CreatedAt:    order.CreatedAt.Unix(),
		UpdatedAt:    order.UpdatedAt.Unix(),
		UserId:       order.UserID,
		ProductId:    order.ProductID,
		BossId:       order.BossID,
		Num:          order.Num,
		AddressName:  address.Name,
		AddressPhone: address.Phone,
		Address:      address.Address,
		Type:         order.Type,
		ProductName:  product.Name,
		ImgPath:      conf.PhotoHost + conf.HttpPort + conf.ProductPhotoPath + product.ImgPath,
		Money:        order.Money,
	}
}

func BuildOrders(ctx context.Context, items []*model.Order) (orders []Order) {
	productDao := dao.NewProductDao(ctx)
	addressDao := dao.NewAddressDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductID)
		if err != nil {
			continue
		}
		address, err := addressDao.GetAddressByAid(item.AddressID)
		if err != nil {
			continue
		}
		order := BuildOrder(item, address, product)
		orders = append(orders, order)
	}

	return orders
}
