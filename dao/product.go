/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-25 10:50:16
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-25 10:54:15
 * @FilePath: /go-mall/dao/product.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"go-mall/model"

	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// 创建商品
func (dao *ProductDao) CreateProduct(product *model.Product) (err error) {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}

// GetProductById 根据id获取Product
func (dao *ProductDao) GetProductById(id uint) (Product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("id = ?", id).First(&Product).Error
	return
}
