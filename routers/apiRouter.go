package routers

import (
	"github.com/gin-gonic/gin"
	"leastMall_gin/controllers/api"
)

var Router = gin.Default()

func init() {
	ns := Router.Group("/api/v1")
	{
		ns.GET("/", api.Get)
		ns.GET("/menu", api.Menu)
	}
}
