package frontend

import (
	"github.com/gin-gonic/gin"
	"leastMall_gin/conn"
	"leastMall_gin/models"
	"net/http"
)

func AddAddress(c *gin.Context) {
	user := models.User{}
	models.Cookie.Get(c, "userinfo", &user)
	name := c.GetString("name")
	phone := c.GetString("phone")
	address := c.GetString("address")
	zipcode := c.GetString("zipcode")
	var addressCount int
	conn.Db.Where("uid=?", user.Id).Table("address").Count(&addressCount)
	if addressCount > 10 {
		data := map[string]interface{}{
			"success": false,
			"message": "增加收货地址失败，收货地址数量超过限制",
		}
		c.JSON(http.StatusOK, data)
		return
	}
	conn.Db.Table("address").Where("uid=?", user.Id).Updates(map[string]interface{}{"default_address": 0})
	addressResult := models.Address{
		Uid:            user.Id,
		Name:           name,
		Phone:          phone,
		Address:        address,
		Zipcode:        zipcode,
		DefaultAddress: 1,
	}
	conn.Db.Create(&addressResult)
	allAddressResult := []models.Address{}
	conn.Db.Where("uid=?", user.Id).Find(&allAddressResult)
	data := map[string]interface{}{
		"success": true,
		"result":  allAddressResult,
	}
	c.JSON(http.StatusOK, data)
}

func GetOneAddressList(c *gin.Context) {
	addressId, err := c.Get("address_id")
	if !err {
		data := map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
		c.JSON(http.StatusOK, data)
		return
	}
	address := models.Address{}
	conn.Db.Where("id=?", addressId).Find(&address)
	data := map[string]interface{}{
		"success": true,
		"result":  address,
	}
	c.JSON(http.StatusOK, data)
}

func GoEditAddressList(c *gin.Context) {
	user := models.User{}
	models.Cookie.Get(c, "userinfo", &user)
	addressId, err := c.Get("address_id")
	if !err {
		data := map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
		c.JSON(http.StatusOK, data)
		return
	}
	name := c.GetString("name")
	phone := c.GetString("phone")
	address := c.GetString("address")
	zipcode := c.GetString("zipcode")
	conn.Db.Table("address").Where("uid=?", user.Id).Updates(map[string]interface{}{"default_address": 0})
	addressModel := models.Address{}
	conn.Db.Where("id=?", addressId).Find(&addressModel)
	addressModel.Name = name
	addressModel.Phone = phone
	addressModel.Address = address
	addressModel.Zipcode = zipcode
	addressModel.DefaultAddress = 1
	conn.Db.Save(&addressModel)
	// 查询当前用户的所有收货地址并返回
	var allAddressResult []models.Address
	conn.Db.Where("uid=?", user.Id).Order("default_address desc").Find(&allAddressResult)

	data := map[string]interface{}{
		"success": true,
		"result":  allAddressResult,
	}
	c.JSON(http.StatusOK, data)

}

func ChangeDefaultAddress(c *gin.Context) {
	user := models.User{}
	models.Cookie.Get(c, "userinfo", &user)
	addressId, err := c.Get("address_id")
	if !err {
		data := map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
		c.JSON(http.StatusOK, data)
		return
	}
	conn.Db.Table("address").Where("uid=?", user.Id).Updates(map[string]interface{}{"default_address": 0})
	conn.Db.Table("address").Where("id=?", addressId).Updates(map[string]interface{}{"default_address": 1})
	data := map[string]interface{}{
		"success": true,
		"result":  "更新默认收获地址成功",
	}
	c.JSON(http.StatusOK, data)
}
