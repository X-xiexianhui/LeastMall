//Package backend
/*
   @author:xie
   @date:2022/6/1
   @note:
*/
package backend

import (
	"github.com/gin-gonic/gin"
	"leastMall_gin/common"
	"leastMall_gin/conn"
	"leastMall_gin/models"
	"strconv"
)

func GetProducts(c *gin.Context) {
	var product []models.Product
	err := conn.Db.Table("product").Find(&product).Error
	if err != nil {
		c.JSON(500, models.NewResponse(false, "获取商品列表失败", "原因："))
	}
	c.JSON(200, models.NewResponse(true, product, "获取商品列表成功"))
}

func AddProduct(c *gin.Context) {
	productName := c.PostForm("product_name")
	price, _ := strconv.ParseFloat(c.PostForm("price"), 64)
	descriptions := c.PostForm("descriptions")
	//商品封面
	file, _ := c.FormFile("cover")
	cover := common.FormatBase64(file)
	product := models.Product{
		ProductName:  productName,
		Price:        price,
		Descriptions: descriptions,
		Cover:        cover,
	}
	//商品相册
	form, _ := c.MultipartForm()
	img := form.File["images"]
	var images []models.Image
	conn.Db.Begin()
	conn.Db.Table("product").Create(&product)
	for _, img := range img {
		image := common.FormatBase64(img)
		images = append(images, models.Image{
			ProductId: product.ProductId,
			Image:     image,
		})
	}
	conn.Db.Table("product").Create(images)
	conn.Db.Commit()
}
