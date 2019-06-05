package main

import (
	"github.com/ejin66/goejin/system"
)

type HomeController struct {
	system.BaseController
}

func (this *HomeController) Filter() (bool, string) {
	this.Ctx.AddHeader("Access-Control-Allow-Origin", "*")
	return true, ""
}

/*              以下方法为自定义方法             */

func (this *HomeController) Index() {
	this.Ctx.SessionStart()
	body := this.Ctx.Body()
	this.Ctx.Response("index....." + body)
}
