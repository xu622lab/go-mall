package v1

import (
	"go-mall/pkg/util"
	"go-mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["file"]
	claim, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	var createProductService service.ProuctService
	if err := ctx.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(ctx.Request.Context(), claim.ID, files)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func ListProduct(ctx *gin.Context) {
	var listProductService service.ProuctService
	if err := ctx.ShouldBind(&listProductService); err == nil {
		res := listProductService.List(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func SearchProduct(ctx *gin.Context) {
	var searchProductService service.ProuctService
	if err := ctx.ShouldBind(&searchProductService); err == nil {
		res := searchProductService.Search(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func ShowProduct(ctx *gin.Context) {
	var showProductService service.ProuctService
	if err := ctx.ShouldBind(&showProductService); err == nil {
		res := showProductService.Show(ctx.Request.Context(), ctx.Param("id"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusOK, ErrorResponse(err))
	}
}
