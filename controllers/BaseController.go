package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

type ReturnMsg struct {
	Code int
	Msg  string
	Data interface{}
}

func (c *BaseController) SuccessJson(data interface{}) {

	res := ReturnMsg{
		200, "success", data,
	}
	c.Data["json"] = res
	c.ServeJSON() //对json进行序列化输出
	c.StopRun()
}

func (c *BaseController) ErrorJson(code int, msg string, data interface{}) {

	res := ReturnMsg{
		code, msg, data,
	}

	c.Data["json"] = res
	c.ServeJSON() //对json进行序列化输出
	c.StopRun()
}
