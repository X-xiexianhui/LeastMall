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
		json := map[string]interface{}{
			"success": false,
			"message": "增加收货地址失败，收货地址数量超过限制",
		}
		c.JSON(http.StatusOK, json)
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
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"result":  allAddressResult,
	}
	c.ServeJSON()
}

func (c *AddressController) GetOneAddressList() {
	addressId, err := c.GetInt("address_id")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
		c.ServeJSON()
		return
	}
	address := models.Address{}
	conn.Db.Where("id=?", addressId).Find(&address)
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"result":  address,
	}
	c.ServeJSON()
}

func (c *AddressController) GoEditAddressList() {
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	addressId, err := c.GetInt("address_id")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
		c.ServeJSON()
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

	c.Data["json"] = map[string]interface{}{
		"success": true,
		"result":  allAddressResult,
	}
	c.ServeJSON()

}

func (c *AddressController) ChangeDefaultAddress() {
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	addressId, err := c.GetInt("address_id")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "传入参数错误",
		}
		c.ServeJSON()
		return
	}
	conn.Db.Table("address").Where("uid=?", user.Id).Updates(map[string]interface{}{"default_address": 0})
	conn.Db.Table("address").Where("id=?", addressId).Updates(map[string]interface{}{"default_address": 1})
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"result":  "更新默认收获地址成功",
	}
	c.ServeJSON()
}
