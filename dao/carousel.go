/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-18 17:07:22
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-21 16:54:00
 * @FilePath: /go-mall/dao/Carousel.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"go-mall/model"

	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

// GetCarouselById 根据id获取Carousel
func (dao *CarouselDao) ListCarousel() (carousel []model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return
}
