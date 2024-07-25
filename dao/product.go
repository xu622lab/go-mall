/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-25 10:50:16
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-25 17:25:36
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

// 查询数据库中符合 condition 条件的 Product 记录的数量
func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	// page.PageNum 是当前页码
	// Offset((page.PageNum - 1) * (page.PageSize)) 计算查询的起始位置，确保从正确的页码开始。
	err = dao.DB.Where(condition).Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).Find(&products).Error
	return
}
