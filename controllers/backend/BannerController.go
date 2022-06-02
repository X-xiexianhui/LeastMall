//Package backend
/*
   @author:xie
   @date:2022/6/1
   @note:
*/
package backend

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"leastMall_gin/conn"
	"leastMall_gin/factory/backendFactory"
	"leastMall_gin/models"
	"strconv"
)

func GetBanner(c *gin.Context) {
	var banner []models.Image
	conn.Db.Table("banner").Find(&banner)
	c.JSON(200, models.NewResponse(true, banner, "查询轮播图"))
}

func AddBanner(c *gin.Context) {
	productId, _ := strconv.ParseInt(c.PostForm("product_id"), 10, 32)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, models.NewResponse(false, "上传图片失败", "原因："+err.Error()))
	}
	picture, err := file.Open()
	if err != nil {
		c.JSON(500, models.NewResponse(false, "上传图片失败", "原因："+err.Error()))
	}
	data, err := ioutil.ReadAll(picture)
	base64Str := base64.StdEncoding.EncodeToString(data)
	banner := models.Image{
		ProductId: productId,
		Image:     base64Str,
	}
	if err := conn.Db.Table("image").Create(&banner).Error; err != nil {
		c.JSON(500, models.NewResponse(false, "图片写入数据库失败", "原因："+err.Error()))
	}
	c.JSON(200, models.NewResponse(true, "上传图片成功", "成功"))
}
func DeleteBanner(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 32)
	banner := backendFactory.SimpleFactory("banner")
	if err := conn.Db.Where("id=?", id).Delete(&banner).Error; err != nil {
		c.JSON(500, models.NewResponse(false, "删除图片失败", "原因："+err.Error()))
	}
	c.JSON(200, models.NewResponse(true, "删除图片成功", "成功"))
}

func QueryBanner(c *gin.Context) {
	productId, _ := strconv.ParseInt(c.Query("product_id"), 10, 32)
	banner := models.Image{}
	if err := conn.Db.Where("product_id=?", productId).First(&banner).Error; err != nil {
		c.JSON(500, models.NewResponse(false, "查询轮播图失败", "原因:"+err.Error()))
	}
	c.JSON(200, models.NewResponse(true, banner, "查询成功"))
}

func UpdateBanner(c *gin.Context) {
	updateColumn := c.PostFormMap("updateColumn")
	banner := models.Image{}
	if err := conn.Db.Model(&banner).Updates(updateColumn).Error; err != nil {
		c.JSON(500, models.NewResponse(false, "修改轮播图失败", "原因："+err.Error()))
	}
	c.JSON(200, models.NewResponse(true, "修改轮播图成功", "操作成功"))
}
