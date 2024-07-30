/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-29 20:39:24
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-30 15:31:11
 * @FilePath: /go-mall/api/v1/pay.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"go-mall/pkg/util"
	"go-mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OrderPay(ctx *gin.Context) {
	orderPay := service.OrderPay{}
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&orderPay); err == nil {
		res := orderPay.PayDown(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}
