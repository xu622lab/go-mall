package service

import "go-mall/model"

type ProuctService struct {
	Id            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"categort_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}
