package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	//c.Ctx.WriteString("api v2")
	c.String(http.StatusOK, "api v2")
}
