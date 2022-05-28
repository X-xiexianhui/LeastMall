//Package frontend
/*
   @author:xie
   @date:2022/5/28
   @note:
*/
package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorldController(c *gin.Context) {
	c.JSON(http.StatusOK, "helloWorld")
}
