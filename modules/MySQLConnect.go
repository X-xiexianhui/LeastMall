//Package modules
/*
   @author:xie
   @date:2022/5/24
   @note:
*/
package modules

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"leastMall_gin/conf"
	"log"
)

var Db *gorm.DB
var err error

func init() {
	mysql := conf.Conf.Mysql
	user := mysql.User
	password := mysql.Password
	host := mysql.Host
	port := mysql.Port
	url := user + ":" + password + "@tcp(" + host + port + ")/shop?charset=utf8&parseTime=True&loc=Local"
	Db, err = gorm.Open("mysql", url)
	if err != nil {
		panic("数据库启动失败……")
	}
	log.Println("数据库启动成功……")
}
