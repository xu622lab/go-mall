/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-26 10:31:59
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-26 10:39:57
 * @FilePath: /go-mall/serializer/product_img.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import (
	"go-mall/conf"
	"go-mall/model"
)

type ProductImg struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductId: item.ProductID,
		ImgPath:   conf.PhotoHost + conf.HttpPort + conf.ProductPhotoPath + item.ImgPath,
	}
}

func BuildProductImgs(items []*model.ProductImg) (productImg []ProductImg) {
	for _, item := range items {
		product := BuildProductImg(item)
		productImg = append(productImg, product)
	}
	return
}
