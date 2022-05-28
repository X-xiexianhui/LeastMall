//Package routers
/*
   @author:xie
   @date:2022/5/27
   @note:
*/
package routers

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func init() {
	Router = gin.Default()
}
