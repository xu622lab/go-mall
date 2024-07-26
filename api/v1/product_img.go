/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-26 10:24:42
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-26 10:26:03
 * @FilePath: /go-mall/api/v1/product_img.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"go-mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProductImg(ctx *gin.Context) {
	var listProductService service.ListProductImg
	if err := ctx.ShouldBind(&listProductService); err == nil {
		res := listProductService.List(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}
