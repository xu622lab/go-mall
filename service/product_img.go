/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-26 10:25:25
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-26 11:27:24
 * @FilePath: /go-mall/service/product_img.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"go-mall/dao"
	"go-mall/serializer"
	"strconv"
)

type ListProductImg struct {
}

func (service *ListProductImg) List(ctx context.Context, pId string) serializer.Response {
	productImgDao := dao.NewProductImgDao(ctx)
	productId, _ := strconv.Atoi(pId)
	productImgs, _ := productImgDao.ListProductImg(uint(productId))
	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(len(productImgs)))
}
