/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-26 15:37:33
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-26 15:47:14
 * @FilePath: /go-mall/api/v1/favorites.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"go-mall/pkg/util"
	"go-mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFavorites(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var createFavoritesService service.FavoriteService
	if err := ctx.ShouldBind(&createFavoritesService); err == nil {
		res := createFavoritesService.Create(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func ListFavorites(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var listFavoritesService service.FavoriteService
	if err := ctx.ShouldBind(&listFavoritesService); err == nil {
		res := listFavoritesService.List(ctx.Request.Context(), claim.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func DeleteFavorites(ctx *gin.Context) {
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var deleteFavoritesService service.FavoriteService
	if err := ctx.ShouldBind(&deleteFavoritesService); err == nil {
		res := deleteFavoritesService.Delete(ctx.Request.Context(), claim.ID, ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}
