package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world!")
}

func (this *MainController) Post() {
	mystruct := `{ "message":"ok" }`
	this.Data["json"] = &mystruct
	this.ServeJSON()
}
