// Package conn Package
/*
   @author:xie
   @date:2022/5/24
   @note:
*/
package conn

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"leastMall_gin/modules"
	"log"
)

var Db *gorm.DB
var err error

func init() {
	mysql := modules.Conf.Mysql
	fmt.Println(mysql)
	user := mysql.User
	password := mysql.Password
	host := mysql.Host
	port := mysql.Port
	url := user + ":" + password + "@tcp(" + host + ":" + port + ")/shop?charset=utf8&parseTime=True&loc=Local"
	Db, err = gorm.Open("mysql", url)
	if err != nil {
		log.Panicln("数据库连接失败……")
	}
	_, _ = fmt.Fprintln(gin.DefaultWriter, "数据库连接成功")
}
