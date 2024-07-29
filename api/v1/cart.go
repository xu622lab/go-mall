/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-26 15:37:33
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-29 15:49:33
 * @FilePath: /go-mall/api/v1/Cart.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"go-mall/pkg/util"
	"go-mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCart(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var createCartService service.CartService
	if err := ctx.ShouldBind(&createCartService); err == nil {
		res := createCartService.Create(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func ListCart(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var listCartService service.CartService
	if err := ctx.ShouldBind(&listCartService); err == nil {
		res := listCartService.List(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func DeleteCart(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var deleteCartService service.CartService
	if err := ctx.ShouldBind(&deleteCartService); err == nil {
		res := deleteCartService.Delete(ctx.Request.Context(), claim.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func UpdateCart(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var updateCartService service.CartService
	if err := ctx.ShouldBind(&updateCartService); err == nil {
		res := updateCartService.Update(ctx.Request.Context(), claim.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}
