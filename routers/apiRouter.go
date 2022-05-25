package routers

import (
	"github.com/gin-gonic/gin"
	"leastMall_gin/controllers/api"
)

var Router = gin.Default()

func init() {
	Router.LoadHTMLGlob("./views/**/**/*")
	ns := Router.Group("/api/v1")
	{
		ns.GET("/", api.GetV1)
		ns.GET("/menu", api.Menu)
	}
}
