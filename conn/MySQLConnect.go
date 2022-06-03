// Package conn Package
/*
   @author:xie
   @date:2022/5/24
   @note:
*/
package conn

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB
var err error

func init() {
	mysqlConfig := Conf.Mysql
	user := mysqlConfig.User
	password := mysqlConfig.Password
	host := mysqlConfig.Host
	port := mysqlConfig.Port
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/shop?charset=utf8&parseTime=True&loc=Local", user, password, host, port)
	Db, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Panicln("数据库连接失败……")
	}
	log.Println("数据库连接成功……")
}
