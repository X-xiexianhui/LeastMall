package frontend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"leastMall_gin/conn"
	"leastMall_gin/models"
	"net/http"
	"net/url"
	"strings"
)

func BaseInit(c *gin.Context) {
	//获取顶部导航
	var topMenu []models.Menu
	if hasTopMenu := models.CacheDb.Get(c, "topMenu", &topMenu); hasTopMenu == true {
		c.JSON(http.StatusOK, topMenu)
	} else {
		conn.Db.Where("status=1 AND position=1").Order("sort desc").Find(&topMenu)
		c.JSON(http.StatusOK, topMenu)
		models.CacheDb.Set(c, "topMenu", topMenu, 30*24*60*60)
	}

	//左侧分类（预加载）
	var productCate []models.ProductCate

	if hasProductCate := models.CacheDb.Get(c, "productCate",
		&productCate); hasProductCate == true {
		c.JSON(http.StatusOK, productCate)
	} else {
		conn.Db.Preload("ProductCateItem",
			func(db *gorm.DB) *gorm.DB {
				return db.Where("product_cate.status=1").
					Order("product_cate.sort DESC")
			}).Where("pid=0 AND status=1").Order("sort desc", true).
			Find(&productCate)
		c.JSON(http.StatusOK, productCate)
		models.CacheDb.Set(c, "productCate", productCate, 30*24*60*60)
	}

	//获取中间导航的数据
	var middleMenu []models.Menu
	if hasMiddleMenu := models.CacheDb.Get(c, "middleMenu",
		&middleMenu); hasMiddleMenu == true {
		c.JSON(http.StatusOK, middleMenu)
	} else {
		conn.Db.Where("status=1 AND position=2").Order("sort desc").
			Find(&middleMenu)

		for i := 0; i < len(middleMenu); i++ {
			//获取关联商品
			middleMenu[i].Relation = strings.ReplaceAll(middleMenu[i].Relation, "，", ",")
			relation := strings.Split(middleMenu[i].Relation, ",")
			var product []models.Product
			conn.Db.Where("id in (?)", relation).Limit(6).Order("sort ASC").
				Select("id,title,product_img,price").Find(&product)
			middleMenu[i].ProductItem = product
		}
		c.JSON(http.StatusOK, middleMenu)
		models.CacheDb.Set(c, "middleMenu", middleMenu, 30*24*60*60)
	}

	//判断用户是否登陆
	user := &models.User{}
	models.Cookie.Get(c, "userinfo", user)
	if len(user.Phone) == 11 {
		str := fmt.Sprintf(`<ul>
			<li class="userinfo">
				<a href="#">%v</a>

				<i class="i"></i>
				<ol>
					<li><a href="/user">个人中心</a></li>

					<li><a href="#">我的收藏</a></li>

					<li><a href="/auth/loginOut">退出登录</a></li>
				</ol>

			</li>
		</ul> `, user.Phone)
		c.JSON(http.StatusOK, str)
	} else {
		str := fmt.Sprintf(`<ul>
			<li><a href="/auth/login" target="_blank">登录</a></li>
			<li>|</li>
			<li><a href="/auth/registerStep1" target="_blank" >注册</a></li>
		</ul>`)
		c.JSON(http.StatusOK, str)
	}
	urlPath, _ := url.Parse(c.Request.URL.String())
	c.JSON(http.StatusOK, urlPath)
}
