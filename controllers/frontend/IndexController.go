package frontend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"leastMall_gin/conn"
	"leastMall_gin/models"
	"net/http"
	"time"
)

func IndexController(c *gin.Context) {
	//初始化
	BaseInit(c)
	//开始时间
	startTime := time.Now().UnixNano()

	//获取轮播图 注意获取的时候要写地址
	var banner []models.Banner
	if hasBanner := models.CacheDb.Get(c, "banner", &banner); hasBanner == true {
		c.JSON(http.StatusOK, banner)
	} else {
		conn.Db.Where("status=1 AND banner_type=1").Order("sort desc").Find(&banner)
		c.JSON(http.StatusOK, banner)
		models.CacheDb.Set(c, "banner", banner, 7*24*60*60)
	}

	//获取手机商品列表
	var redisPhone []models.Product
	if hasPhone := models.CacheDb.Get(c, "phone", &redisPhone); hasPhone == true {
		c.JSON(http.StatusOK, redisPhone)
	} else {
		phone := models.GetProductByCategory(1, "hot", 8)
		c.JSON(http.StatusOK, redisPhone)
		models.CacheDb.Set(c, "phone", phone, 7*24*60*60)
	}
	//获取电视商品列表
	var redisTv []models.Product
	if hasTv := models.CacheDb.Get(c, "tv", &redisTv); hasTv == true {
		c.JSON(http.StatusOK, redisTv)
	} else {
		tv := models.GetProductByCategory(4, "best", 8)
		c.JSON(http.StatusOK, tv)
		models.CacheDb.Set(c, "tv", tv, 7*24*60*60)
	}

	//结束时间
	endTime := time.Now().UnixNano()

	fmt.Println("执行时间", endTime-startTime)
}
