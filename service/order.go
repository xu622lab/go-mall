/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-29 10:07:17
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-29 20:34:48
 * @FilePath: /go-mall/service/Order.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"fmt"
	"go-mall/dao"
	"go-mall/model"
	"go-mall/pkg/e"
	"go-mall/serializer"
	"math/rand"
	"strconv"
	"time"
)

type OrderService struct {
	ProductId uint    `json:"product_id" form:"product_id"`
	Num       int     `json:"num" form:"num"`
	AddressId uint    `json:"address_id" form:"address_id"`
	Money     float64 `json:"money" form:"money"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	UserId    uint    `json:"user_id" form:"user_id"`
	OrderNum  int     `json:"order_num" form:"order_num"`
	Type      int     `json:"type" form:"type"`
	model.BasePage
}

func (service *OrderService) Create(ctx context.Context, uId uint) serializer.Response {
	var order *model.Order
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order = &model.Order{
		UserID:    uId,
		ProductID: service.ProductId,
		BossID:    service.BossId,
		Num:       service.Num,
		Money:     service.Money,
		Type:      1, // 默认未支付
	}

	// 检验地址是否存在
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(service.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.AddressID = address.ID
	// 订单号创建，自动生成的随机number+唯一表示的product id + 用户的 id
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000))
	productNum := strconv.Itoa(int(service.ProductId))
	userNum := strconv.Itoa(int(service.UserId))
	number = number + productNum + userNum
	orderNum, _ := strconv.ParseUint(number, 10, 64)
	order.OrderNum = orderNum

	err = orderDao.CreateOrder(order)
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

func (service *OrderService) Show(ctx context.Context, uId uint, oId string) serializer.Response {
	orderId, _ := strconv.Atoi(oId)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order, err := orderDao.GetOrderByOid(uint(orderId), uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 获取地址
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(order.AddressID)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 获取商品信息
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(order.ProductID)
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
		Data:   serializer.BuildOrder(order, address, product),
	}
}

func (service *OrderService) Delete(ctx context.Context, uId uint, aId string) serializer.Response {
	orderId, _ := strconv.Atoi(aId)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	err := orderDao.DeleteOrderByAid(uId, uint(orderId))
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

func (service *OrderService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	orderDao := dao.NewOrderDao(ctx)

	condition := make(map[string]interface{})
	if service.Type != 0 {
		condition["type"] = service.Type
	}
	condition["user_id"] = uId
	orderList, total, err := orderDao.ListOrderByCondition(condition, service.BasePage)
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

		Data: serializer.BuildListResponse(serializer.BuildOrders(ctx, orderList), uint(total)),
	}
}
