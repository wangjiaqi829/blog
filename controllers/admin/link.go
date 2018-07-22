package admin

import (
	"blog/models"
	"strings"
)

type LinkController struct {
	BaseController
}

//友链列表
func (this *LinkController) List() {
	var list []*models.Link
	var link models.Link
	link.Query().OrderBy("-rank").All(&list)
	this.Data["list"] = list
	this.display()
}

//添加友链
func (this *LinkController) Add() {
	if this.Ctx.Request.Method == "POST" {
		var link models.Link
		siteName := strings.TrimSpace(this.GetString("siteName"))
		url := strings.TrimSpace(this.GetString("url"))
		rank, _ := this.GetInt64("rank")
		link.SiteName = siteName
		link.Url = url
		link.Rank = int8(rank)
		if err := link.Insert(); err != nil {
			this.showmsg(err.Error())
		}
		this.Redirect("/admin/link/list", 302)

	}
	this.display()
}

//编辑友链
func (this *LinkController) Edit() {
	id, _ := this.GetInt64("id")
	link := models.Link{Id: id}
	if err := link.Read(); err != nil {
		this.showmsg("友链不存在")
	}

	if this.Ctx.Request.Method == "POST" {
		siteName := strings.TrimSpace(this.GetString("siteName"))
		url := strings.TrimSpace(this.GetString("url"))
		rank, _ := this.GetInt64("rank")
		link.SiteName = siteName
		link.Url = url
		link.Rank = int8(rank)
		link.Update()
		this.Redirect("/admin/link/list", 302)
	}
	this.Data["link"] = link
	this.display()
}

//删除友链
func (this *LinkController) Delete() {
	id, _ := this.GetInt64("id")
	link := models.Link{Id: id}
	if link.Read() == nil {
		link.Delete()
	}
	this.Redirect("/admin/link/list", 302)
}
