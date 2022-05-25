package routers

import (
	"leastMall_gin/controllers/api"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSRouter("/", &api.V1Controller{}),
		beego.NSRouter("/menu", &api.V1Controller{}, "get:Menu"),
	)
	beego.AddNamespace(ns)
}
