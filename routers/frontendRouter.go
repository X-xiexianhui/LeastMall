package routers

import "leastMall_gin/controllers/frontend"

func init() {
	Router.GET("/", frontend.HelloWorldController)
	Router.GET("/test", frontend.Test)
}
