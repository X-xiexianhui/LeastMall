//Package backend
/*
   @author:xie
   @date:2022/6/1
   @note:
*/
package backend

import (
	"github.com/gin-gonic/gin"
	"leastMall_gin/conn"
	"leastMall_gin/models"
)

func GetProducts(c *gin.Context) {
	var product []models.Product
	err := conn.Db.Table("product").Find(&product).Error
	if err != nil {
		c.JSON(500, models.NewResponse(false, "获取商品列表失败", "原因："))
	}
	c.JSON(200, models.NewResponse(true, product, "获取商品列表成功"))
}
