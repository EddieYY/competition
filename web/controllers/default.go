package controllers

import (
	"github.com/EddieYY/competition/web/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	beego.Debug("======")
	dt := models.QureyPrice()
	beego.Debug(dt)
	c.TplName = "index.tpl"
}
