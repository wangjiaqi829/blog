package blog

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	page     int
	pagesize int
}

func (this *BaseController) Prepare() {

}

func (this *BaseController) display(tpl string) {

	this.Layout =  "blog/layout.html"
	this.TplName = "blog/" + tpl + ".html"
}

