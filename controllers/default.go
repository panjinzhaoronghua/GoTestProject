package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "你个我"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

}
