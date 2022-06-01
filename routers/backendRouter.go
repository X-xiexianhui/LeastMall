package routers

import "leastMall_gin/controllers/backend"

func init() {
	backGroup := Router.Group("/back")
	{
		backGroup.GET("/banner/get", backend.GetBanner)
		backGroup.POST("/banner/add", backend.AddBanner)
		backGroup.DELETE("/banner/delete", backend.DeleteBanner)
		backGroup.GET("/banner/query", backend.QueryBanner)
		backGroup.POST("/banner/update", backend.UpdateBanner)
	}
}
