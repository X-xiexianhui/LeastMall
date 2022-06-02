//Package routers
/*
   @author:xie
   @date:2022/5/27
   @note:
*/
package routers

import (
	"github.com/gin-gonic/gin"
	"leastMall_gin/controllers"
)

var Router = gin.Default()

func init() {
	//轮播图api
	Router.GET("/banner/get", controllers.GetBanner)
	Router.POST("/banner/add", controllers.AddBanner)
	Router.DELETE("/banner/delete", controllers.DeleteBanner)
	Router.GET("/banner/query", controllers.QueryBanner)
	Router.POST("/banner/update", controllers.UpdateBanner)
	//商品api
	Router.GET("/product/get", controllers.GetProducts)
	Router.POST("/product/add", controllers.AddProduct)
	Router.DELETE("/product/delete", controllers.DeleteProduct)
	Router.POST("/product/update", controllers.UpdateProduct)
	//商品图片api
	Router.GET("/image/get", controllers.GetImage)
	Router.POST("/image/add", controllers.AddImages)
	Router.DELETE("/image/delete", controllers.DeleteImage)
}
