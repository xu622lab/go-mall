/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-29 10:01:13
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-29 15:48:23
 * @FilePath: /go-mall/dao/Cart.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"go-mall/model"

	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}
func (dao *CartDao) CreateCart(in *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Create(&in).Error
}

func (dao *CartDao) GetCartByAid(aId uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id=?", aId).First(&cart).Error
	return
}

func (dao *CartDao) DeleteCartByAid(cId, uId uint) error {
	return dao.DB.Model(&model.Cart{}).Where("id=?", cId).Delete(&model.Cart{}).Error
}

func (dao *CartDao) ListCartByUserId(uId uint) (cart []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id=?", uId).Find(&cart).Error
	return
}

func (dao *CartDao) UpdateCartNumById(cId uint, num int) error {
	return dao.DB.Model(&model.Cart{}).Where("id=?", cId).Update("num", num).Error
}
