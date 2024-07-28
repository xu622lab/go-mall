/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-26 15:08:02
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-26 15:11:30
 * @FilePath: /go-mall/serializer/category.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import "go-mall/model"

type Category struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64  `json:"create_at"`
}

func BuildCategory(item *model.Category) Category {
	return Category{
		Id:           item.ID,
		CategoryName: item.CategoryName,
		CreateAt:     item.CreatedAt.Unix(),
	}
}

func BuildCategorys(items []*model.Category) (categorys []Category) {
	for _, item := range items {
		category := BuildCategory(item)
		categorys = append(categorys, category)
	}

	return categorys
}
