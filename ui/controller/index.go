package controller

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Ctx.WriteString("hello world")
}
