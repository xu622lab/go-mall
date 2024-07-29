/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-29 10:01:13
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-29 11:12:36
 * @FilePath: /go-mall/dao/address.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"go-mall/model"

	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}
func (dao *AddressDao) CreateAddress(in *model.Address) error {
	return dao.DB.Model(&model.Address{}).Create(&in).Error
}

func (dao *AddressDao) GetAddressByAid(aId uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=?", aId).First(&address).Error
	return
}

func (dao *AddressDao) DeleteAddressByAid(aId, uId uint) error {
	return dao.DB.Model(&model.Address{}).Where("id=?", aId).Delete(&model.Address{}).Error
}

func (dao *AddressDao) ListAddressByUserId(uId uint) (address []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("user_id=?", uId).Find(&address).Error
	return
}

func (dao *AddressDao) UpdateAddressByUserId(aId uint, address *model.Address) error {
	return dao.DB.Model(&model.Address{}).Where("id=?", aId).Updates(&address).Error
}
