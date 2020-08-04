package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

type ReturnMsg struct {
	Code int
	Msg  string
	Data interface{}
}

func (this *BaseController) SuccessJson(data interface{}) {

	res := ReturnMsg{
		200, "success", data,
	}
	this.Data["json"] = res
	this.ServeJSON() //对json进行序列化输出
	this.StopRun()
}

func (this *BaseController) ErrorJson(code int, msg string, data interface{}) {

	res := ReturnMsg{
		code, msg, data,
	}

	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}
