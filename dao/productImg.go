/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-25 10:50:16
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-26 11:27:47
 * @FilePath: /go-mall/dao/ProductImg.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"go-mall/model"

	"gorm.io/gorm"
)

// 商品展示图，在内容详情显示
type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

// 创建
func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) (err error) {
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
}

// ListProductImg 根据id获取ProductImg
func (dao *ProductImgDao) ListProductImg(id uint) (ProductImg []*model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).Where("product_id = ?", id).First(&ProductImg).Error
	return
}
