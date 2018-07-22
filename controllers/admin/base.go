package admin

import (
	"github.com/astaxie/beego"
	"strings"
)

const (
	BIG_PIC_PATH   = "./static/upload/bigpic/"
	SMALL_PIC_PATH = "./static/upload/smallpic/"
	FILE_PATH      = "./static/upload/attachment/"
)

var pathArr []string = []string{"", BIG_PIC_PATH, SMALL_PIC_PATH, FILE_PATH}

type BaseController struct {
	beego.Controller
	moduleName     string
	controllerName string
	actionName     string
}

func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.moduleName = "admin"
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.checkPermission()
}

//登录状态验证
/*func (this *BaseController) auth() {
	if this.controllerName == "account" && (this.actionName == "login" || this.actionName == "logout") {

	} else {
		arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
		if len(arr) == 2 {
			idstr, password := arr[0], arr[1]
			userid, _ := strconv.ParseInt(idstr, 10, 0)
			if userid > 0 {
				var user models.User
				user.Id = userid
				if user.Read() == nil && password == models.Md5([]byte(this.getClientIp()+"|"+user.Password)) {
					this.userid = user.Id
					this.username = user.Username
				}
			}
		}
		if this.userid == 0 {
			this.Redirect("/admin/login", 302)
		}
	}
}*/

//渲染模版
func (this *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) == 1 {
		tplname = this.moduleName + "/" + tpl[0] + ".html"
	} else {
		tplname = this.moduleName + "/" + this.controllerName + "_" + this.actionName + ".html"
	}
	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.Layout = this.moduleName + "/layout.html"
	this.TplName = tplname
}

//显示错误提示
func (this *BaseController) showmsg(msg ...string) {
	if len(msg) == 1 {
		msg = append(msg, this.Ctx.Request.Referer())
	}
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username
	this.Data["msg"] = msg[0]
	this.Data["redirect"] = msg[1]
	this.Layout = this.moduleName + "/layout.html"
	this.TplName = this.moduleName + "/" + "showmsg.html"
	this.Render()
	this.StopRun()
}

//是否post提交
func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

//权限验证
func (this *BaseController) checkPermission() {
	if this.userid != 1 && this.controllerName == "user" {
		this.showmsg("抱歉，只有超级管理员才能进行该操作！")
	}
}


