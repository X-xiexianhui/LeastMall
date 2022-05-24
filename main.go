//Package leastMall_gin
/*
   @author:xie
   @date:2022/5/23
   @note:
*/
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//启动gin
	r := gin.Default()
	_ = r.Run(":8080")
}
