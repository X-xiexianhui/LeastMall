// Package controllers
/*
   @author:xie
   @date:2022/6/1
   @note:
*/
package controllers

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
	conn.Db.Table("product").Create(&product)
	c.JSON(200, models.NewResponse(true, "添加商品成功", "操作成功"))
}

func DeleteProduct(c *gin.Context) {
	id := c.Query("id")
	err := conn.Db.Table("product").Delete(&models.Product{}, id).Error
	if err != nil {
		c.JSON(500, models.NewResponse(false, "删除商品失败", "原因："+err.Error()))
	}
	c.JSON(200, models.NewResponse(true, "删除成功", "操作成功"))
}

func UpdateProduct(c *gin.Context) {
	id := c.PostForm("id")
	update := c.PostFormMap("update")
	file, err := c.FormFile("cover")
	if err != nil {
		cover := common.FormatBase64(file)
		update["cover"] = cover
	}
	conn.Db.Table("product").Where("id=?", id).Updates(update)
}

func AddImages(c *gin.Context) {
	//商品相册
	productId, _ := strconv.ParseInt(c.PostForm("product_id"), 10, 64)
	form, _ := c.MultipartForm()
	img := form.File["images"]
	var images []models.Image
	for _, img := range img {
		image := common.FormatBase64(img)
		images = append(images, models.Image{
			ProductId: productId,
			Image:     image,
		})
	}
	conn.Db.Table("images").Create(&images)
	c.JSON(200, models.NewResponse(true, "添加图片成功", "操作成功"))
}

func GetImage(c *gin.Context) {
	productId, _ := strconv.ParseInt(c.PostForm("product_id"), 10, 64)
	var images []models.Image
	err := conn.Db.Table("images").Where("product_id=?", productId).Find(&images).Error
	if err != nil {
		c.JSON(500, models.NewResponse(false, images, "获取商品图片失败"))
	}
	c.JSON(200, models.NewResponse(true, images, "获取商品图片成功"))
}

func DeleteImage(c *gin.Context) {
	id := c.PostForm("id")
	err := conn.Db.Table("images").Delete(&models.Image{}, id).Error
	if err != nil {
		c.JSON(500, models.NewResponse(false, "删除图片失败", "操作失败"))
	}
	c.JSON(200, models.NewResponse(true, "删除图片成功", "操作成功"))
}
