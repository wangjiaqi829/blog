package main

import (
	"github.com/astaxie/beego"
	_"blog/routers"
)

func main() {

	//日志的保存位置以及日志文件的名称
	beego.SetLogger("file", `{"filename":"logs/blog.log","daily":true}`)

	//输出日志级别  为了方便测试，先将日志级别设为最低
	beego.SetLevel(beego.LevelInformational)

	//默认不输出文件名和行号，为了方便测试，现将行号以及文件名输出
	beego.SetLogFuncCall(true)

	//运行程序
	beego.Run()
}


