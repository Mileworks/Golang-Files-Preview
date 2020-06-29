package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["plm"] = "PLM 文件预览服务"
	c.TplName = "index.tpl"
}
