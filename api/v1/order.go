/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-26 15:37:33
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-29 15:49:33
 * @FilePath: /go-mall/api/v1/Order.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"go-mall/pkg/util"
	"go-mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var createOrderService service.OrderService
	if err := ctx.ShouldBind(&createOrderService); err == nil {
		res := createOrderService.Create(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func ListOrder(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var listOrderService service.OrderService
	if err := ctx.ShouldBind(&listOrderService); err == nil {
		res := listOrderService.List(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func DeleteOrder(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var deleteOrderService service.OrderService
	if err := ctx.ShouldBind(&deleteOrderService); err == nil {
		res := deleteOrderService.Delete(ctx.Request.Context(), claim.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func ShowOrder(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var showOrderService service.OrderService
	if err := ctx.ShouldBind(&showOrderService); err == nil {
		res := showOrderService.Show(ctx.Request.Context(), claim.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}
