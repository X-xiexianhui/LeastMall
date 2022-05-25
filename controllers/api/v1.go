package api

import (
	"leastMall_gin/conn"
	"leastMall_gin/models"
)

type V1Controller struct {
}

func (c *V1Controller) Get() {
	c.Ctx.WriteString("api v1")
}

func (c *V1Controller) Menu() {
	menu := []models.Menu{}
	conn.Db.Find(&menu)
	c.Data["json"] = menu
	c.ServeJSON()
}
