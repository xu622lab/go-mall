/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-29 10:01:13
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-29 20:34:11
 * @FilePath: /go-mall/dao/Order.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"go-mall/model"

	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}
func (dao *OrderDao) CreateOrder(in *model.Order) error {
	return dao.DB.Model(&model.Order{}).Create(&in).Error
}

func (dao *OrderDao) GetOrderByOid(id, uId uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id=? AND user_id=?", id, uId).First(&order).Error
	return
}

func (dao *OrderDao) DeleteOrderByAid(aId, uId uint) error {
	return dao.DB.Model(&model.Order{}).Where("id=?", aId).Delete(&model.Order{}).Error
}

func (dao *OrderDao) ListOrderByUserId(uId uint) (order []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("user_id=?", uId).Find(&order).Error
	return
}

func (dao *OrderDao) UpdateOrderByUserId(aId uint, order *model.Order) error {
	return dao.DB.Model(&model.Order{}).Where("id=?", aId).Updates(&order).Error
}

func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (order []*model.Order, total int64, err error) {
	err = dao.DB.Model(&model.Order{}).Where(condition).Count(&total).Error
	err = dao.DB.Model(&model.Order{}).
		Where(condition).
		Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).
		Find(&order).Error
	return
}
