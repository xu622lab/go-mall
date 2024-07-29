/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-29 15:03:04
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-29 15:54:14
 * @FilePath: /go-mall/serializer/cart.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import (
	"context"
	"go-mall/conf"
	"go-mall/dao"
	"go-mall/model"
)

type Cart struct {
	Id            uint   `json:"id"`
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	CreateAt      int64  `json:"created_at"`
	Num           int    `json:"num"`
	Name          string `json:"name"`
	MaxNum        int    `json:"max_num"`
	ImgPath       string `json:"img_path"`
	Check         bool   `json:"check"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.User) Cart {
	return Cart{
		Id:            cart.ID,
		UserId:        cart.UserID,
		ProductId:     cart.ProductID,
		CreateAt:      cart.CreatedAt.Unix(),
		Num:           int(cart.Num),
		Name:          product.Name,
		MaxNum:        int(cart.MaxNum),
		ImgPath:       conf.PhotoHost + conf.HttpPort + conf.ProductPhotoPath + product.ImgPath,
		Check:         cart.Check,
		DiscountPrice: product.DiscountPrice,
		BossId:        boss.ID,
		BossName:      boss.UserName,
	}
}

func BuildCarts(ctx context.Context, items []*model.Cart) (carts []Cart) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductID)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserById(item.UserID)
		if err != nil {
			continue
		}
		cart := BuildCart(item, product, boss)
		carts = append(carts, cart)
	}

	return carts
}
