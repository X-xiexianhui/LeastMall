//Package leastMall_gin
/*
   @author:xie
   @date:2022/5/23
   @note:
*/
package main

import (
	"github.com/gin-gonic/gin"
	"leastMall_gin/conn"
)

func main() {
	//日志着色
	gin.ForceConsoleColor()
	//启动gin
	r := gin.Default()
	conn.Db.LogMode(true)
	_ = r.Run(":8080")
}
