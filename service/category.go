/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-26 15:06:05
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-26 15:17:02
 * @FilePath: /go-mall/service/category.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"go-mall/dao"
	"go-mall/pkg/e"
	"go-mall/pkg/util"
	"go-mall/serializer"
)

type CategoryService struct {
}

func (service *CategoryService) List(ctx context.Context) serializer.Response {
	categoryDao := dao.NewCategoryDao(ctx)
	code := e.Success
	category, err := categoryDao.ListCategory()
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("err", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildCategorys(category), uint(len(category)))
}
