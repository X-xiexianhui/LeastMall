//Package leastMall_gin
/*
   @author:xie
   @date:2022/5/23
   @note:
*/
package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	//日志着色
	gin.ForceConsoleColor()
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	//启动gin
	r := gin.Default()
	_ = r.Run(":8080")
}
