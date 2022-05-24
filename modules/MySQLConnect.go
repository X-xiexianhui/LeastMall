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
)

var Db *gorm.DB
var err error

func init() {

}
