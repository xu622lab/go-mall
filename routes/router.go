/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-16 15:41:08
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-25 16:20:23
 * @FilePath: /go-mall/routes/router.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-16 15:41:08
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-21 20:02:54
 * @FilePath: /go-mall/routes/router.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @Author: xuzhaoyang 15809246338@163.com
 * @Date: 2024-07-16 15:41:08
 * @LastEditors: xuzhaoyang 15809246338@163.com
 * @LastEditTime: 2024-07-21 15:51:31
 * @FilePath: /go-mall/routes/router.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package routes

import (
	api "go-mall/api/v1"
	"go-mall/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "success")
		})

		// 用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		// 轮播图
		v1.GET("carousels", api.ListCarousel)

		// 商品获取操作
		v1.GET("products", api.ListProduct)

		authed := v1.Group("/") // 需要登录保护
		authed.Use(middleware.JWT())
		{
			// 用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)

			// 显示金额
			authed.POST("money", api.ShowMoney)

			// 商品操作
			authed.POST("product", api.CreateProduct)
		}

	}

	return r
}
