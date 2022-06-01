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
	var image []string
	conn.Db.Table("product").Model(product).Related(image, "images")
}
