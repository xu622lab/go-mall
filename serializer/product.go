/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-25 15:20:07
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-25 16:55:15
 * @FilePath: /go-mall/serializer/product.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import (
	"go-mall/conf"
	"go-mall/model"
)

type Product struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint64 `json:"view"`
	CreatedAt     int64  `json:"created_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossID        int    `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

// 序列化商品
func BuildProduct(item *model.Product) Product {
	p := Product{
		ID:            item.ID,
		Name:          item.Name,
		CategoryID:    item.CategoryID,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       conf.PhotoHost + conf.HttpPort + conf.ProductPhotoPath + item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		View:          item.View(),
		Num:           item.Num,
		OnSale:        item.OnSale,
		CreatedAt:     item.CreatedAt.Unix(),
		BossID:        int(item.BossID),
		BossName:      item.BossName,
		BossAvatar:    conf.PhotoHost + conf.HttpPort + conf.AvatarPath + item.BossAvatar,
	}

	// if conf.UploadModel == consts.UploadModelOss {
	// 	p.ImgPath = item.ImgPath
	// 	p.BossAvatar = item.BossAvatar
	// }

	return p

}

// 分页查询序列化
func BuildProducts(items []*model.Product) (products []Product) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}
	return products
}
