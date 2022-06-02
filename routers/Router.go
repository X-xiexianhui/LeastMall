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
	Router.GET("/banner/get", controllers.GetBanner)
	Router.POST("/banner/add", controllers.AddBanner)
	Router.DELETE("/banner/delete", controllers.DeleteBanner)
	Router.GET("/banner/query", controllers.QueryBanner)
	Router.POST("/banner/update", controllers.UpdateBanner)

	Router.GET("/product/get", controllers.GetProducts)
	Router.POST("/product/add", controllers.AddProduct)
}
