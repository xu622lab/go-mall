/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-29 20:44:06
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-30 15:31:08
 * @FilePath: /go-mall/service/pay.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"errors"
	"fmt"
	"go-mall/dao"
	"go-mall/model"
	"go-mall/pkg/e"
	"go-mall/pkg/util"
	"go-mall/serializer"
	"strconv"
)

type OrderPay struct {
	OrderId   uint    `json:"order_id" form:"order_id"`
	Money     float64 `json:"money" form:"money"`
	OrderNo   string  `json:"order_no" form:"order_no"`
	ProductId uint    `json:"product_id" form:"product_id"`
	PayTime   string  `json:"pay_time" form:"pay_time"`
	Sign      string  `json:"sign" form:"sign"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	Num       int     `josn:"num" form:"num"`
	Key       string  `json:"key" form:"key"` // 支付的金额
}

func (service *OrderPay) PayDown(ctx context.Context, uId uint) serializer.Response {
	util.Encrypt.SetKey(service.Key)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	// 使用事务
	tx := orderDao.Begin()
	order, err := orderDao.GetOrderByOid(service.OrderId, uId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	money := order.Money
	num := order.Money
	money = money * float64(num)

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 对钱进行解密，减去订单，再加密保存
	moneyStr := util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)
	if moneyFloat-money < 0.0 {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("金额不足").Error(),
		}
	}

	// 更新账户金额
	finMoney := fmt.Sprintf("%f", moneyFloat-money)
	user.Money = util.Encrypt.AesEncoding(finMoney)

	userDao = dao.NewUserDaoByDB(userDao.DB)
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("用户金额更新失败").Error(),
		}
	}

	// 商家加钱
	var boss *model.User
	boss, err = userDao.GetUserById(service.BossId)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("商家信息获取失败").Error(),
		}
	}
	moneyStr = util.Encrypt.AesDecoding(boss.Money)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)
	finMoney = fmt.Sprintf("%f", moneyFloat+money)
	boss.Money = util.Encrypt.AesEncoding(finMoney)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("商家金额更新失败").Error(),
		}
	}

	// 对应商品数量 - 1
	var product *model.Product
	productDao := dao.NewProductDao(ctx)
	product, err = productDao.GetProductById(service.ProductId)
	product.Num -= int(num)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("商品查询失败").Error(),
		}
	}
	err = productDao.UpdateProduct(service.ProductId, product)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("商品数量更新失败").Error(),
		}
	}

	// 订单删除
	err = orderDao.DeleteOrderByAid(service.OrderId, uId)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("订单删除失败").Error(),
		}
	}

	// 自己的商品+1 同一件商品？title？保留一个原商品id。数据库加一个字段
	productUser := model.Product{
		Name:          product.Name,
		CategoryID:    product.CategoryID,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.Price,
		OnSale:        false,
		Num:           1,
		BossID:        uId,
		BossName:      user.UserName,
		BossAvatar:    user.Avatar,
	}

	err = productDao.CreateProduct(&productUser)
	if err != nil {
		tx.Rollback()
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("商品重新创建失败").Error(),
		}
	}

	tx.Commit()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}
