/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-26 15:37:33
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-29 11:08:03
 * @FilePath: /go-mall/api/v1/Address.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"go-mall/pkg/util"
	"go-mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAddress(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var createAddressService service.AddressService
	if err := ctx.ShouldBind(&createAddressService); err == nil {
		res := createAddressService.Create(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func ListAddress(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var listAddressService service.AddressService
	if err := ctx.ShouldBind(&listAddressService); err == nil {
		res := listAddressService.List(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func DeleteAddress(ctx *gin.Context) {
	var deleteAddressService service.AddressService
	if err := ctx.ShouldBind(&deleteAddressService); err == nil {
		res := deleteAddressService.Delete(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func ShowAddress(ctx *gin.Context) {
	var showAddressService service.AddressService
	if err := ctx.ShouldBind(&showAddressService); err == nil {
		res := showAddressService.Show(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func UpdateAddress(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var updateAddressService service.AddressService
	if err := ctx.ShouldBind(&updateAddressService); err == nil {
		res := updateAddressService.Update(ctx.Request.Context(), claim.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}
