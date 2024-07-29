/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-29 10:07:17
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-29 15:49:26
 * @FilePath: /go-mall/service/Cart.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"go-mall/dao"
	"go-mall/model"
	"go-mall/pkg/e"
	"go-mall/serializer"
	"strconv"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       int  `json:"num" form:"num"`
}

func (service *CartService) Create(ctx context.Context, uId uint) serializer.Response {
	var cart *model.Cart
	code := e.Success

	// 判断有没有这个商品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	cartDao := dao.NewCartDao(ctx)
	cart = &model.Cart{
		UserID:    uId,
		ProductID: service.ProductId,
		BossID:    service.BossId,
	}
	err = cartDao.CreateCart(cart)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserById(service.BossId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

func (service *CartService) Delete(ctx context.Context, uId uint, cId string) serializer.Response {
	cartId, _ := strconv.Atoi(cId)
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	err := cartDao.DeleteCartByAid(uId, uint(cartId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *CartService) Update(ctx context.Context, uId uint, cId string) serializer.Response {
	cartId, _ := strconv.Atoi(cId)
	code := e.Success
	cartDao := dao.NewCartDao(ctx)

	err := cartDao.UpdateCartNumById(uint(cartId), service.Num)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *CartService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	CartDao := dao.NewCartDao(ctx)

	cartList, err := CartDao.ListCartByUserId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarts(ctx, cartList),
	}
}
