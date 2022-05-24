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
	"log"
)

var Db *gorm.DB
var err error

func init() {
	mysql := Conf.Mysql
	user := mysql.User
	password := mysql.Password
	host := mysql.Host
	port := mysql.Port
	url := user + ":" + password + "@tcp(" + host + ":" + port + ")/shop?charset=utf8&parseTime=True&loc=Local"
	Db, err = gorm.Open("mysql", url)
	Db.LogMode(true)
	if err != nil {
		log.Panicln("数据库连接失败……")
	}
	log.Println("数据库连接成功……")
}
