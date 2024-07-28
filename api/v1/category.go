/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-26 15:01:48
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-26 15:17:05
 * @FilePath: /go-mall/api/v1/category.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"go-mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategory(ctx *gin.Context) {
	var listCategory service.CategoryService
	if err := ctx.ShouldBind(&listCategory); err == nil {
		res := listCategory.List(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}
