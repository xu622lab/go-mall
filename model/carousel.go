/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-16 11:26:23
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-16 11:26:37
 * @FilePath: /go-mall/model/carousel.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import "github.com/jinzhu/gorm"

type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductID uint `gorm:"not null"`
}
