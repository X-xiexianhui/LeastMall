package api

import (
	"github.com/gin-gonic/gin"
	"leastMall_gin/conn"
	"leastMall_gin/models"
	"net/http"
)

func Get(c *gin.Context) {
	//c.Ctx.WriteString("api v1")
	c.String(http.StatusOK, "api v1")

}

func Menu(c *gin.Context) {
	var menu []models.Menu
	conn.Db.Find(&menu)
	//c.Data["json"] = menu
	//c.ServeJSON()
	c.JSON(http.StatusOK, menu)
}
