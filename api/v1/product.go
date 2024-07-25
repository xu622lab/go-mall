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
