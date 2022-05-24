//Package leastMall_gin
/*
   @author:xie
   @date:2022/5/23
   @note:
*/
package main

import (
	"github.com/gin-gonic/gin"
	"leastMall_gin/conf"
)

func main() {
	//启动gin
	r := gin.Default()
	//初始化配置
	conf.InitConfig()
	_ = r.Run(":8080")
}
