/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-16 17:01:21
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-21 10:44:41
 * @FilePath: /go-mall/api/v1/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"go-mall/pkg/util"
	"go-mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	var userRegister service.UserService
	if err := ctx.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

func UserLogin(ctx *gin.Context) {
	var userLogin service.UserService
	if err := ctx.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

func UserUpdate(ctx *gin.Context) {
	var userUpdate service.UserService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&userUpdate); err == nil {
		res := userUpdate.Update(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

func UploadAvatar(ctx *gin.Context) {
	file, fileHeader, _ := ctx.Request.FormFile("file")
	fileSize := fileHeader.Size
	var uploadAvatar service.UserService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&uploadAvatar); err == nil {
		res := uploadAvatar.Post(ctx.Request.Context(), claims.ID, file, fileSize)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

func SendEmail(ctx *gin.Context) {
	var sendEmail service.SendEmailService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&sendEmail); err == nil {
		res := sendEmail.Send(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

// ValidEmail 验证邮箱
func ValidEmail(ctx *gin.Context) {
	var validEmail service.ValidEmailService
	if err := ctx.ShouldBind(&validEmail); err == nil {
		res := validEmail.Valid(ctx.Request.Context(), ctx.GetHeader("Authorization"))
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}

// ShowMoney 显示金额
func ShowMoney(ctx *gin.Context) {
	var ShowMoney service.ShowMoneyService
	claims, _ := util.ParseToken(ctx.GetHeader("Authorization"))
	if err := ctx.ShouldBind(&ShowMoney); err == nil {
		res := ShowMoney.Show(ctx.Request.Context(), claims.ID)
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err)
	}
}
