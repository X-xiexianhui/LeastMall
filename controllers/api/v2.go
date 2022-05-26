package api

import (
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	"net/http"
)

type V2Controller struct {
	beego.Controller
}

func GetV2(c *gin.Context) {
	//c.Ctx.WriteString("api v2")
	c.String(http.StatusOK, "api v2")
}
