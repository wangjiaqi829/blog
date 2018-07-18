package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"blog/models"
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


func init() {

	err := beego.LoadAppConfig("ini", "./src/ty/conf/app.conf")
	if err != nil {
		beego.Error("loadAppConfig error:", err.Error())
	}
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}

	beego.Info(dbhost, dbuser)
	dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Local"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dburl)

	//将定义的iot Model 进行注册
	orm.RegisterModel(new(models.Article),new(models.Comment),new(models.Tag),new(models.User))

	// 开启 ORM 调试模式
	orm.Debug = true

	// 自动建表
	orm.RunSyncdb("default", false, true)


	if beego.AppConfig.String("RunMode") == "dev" {
		orm.Debug = false
	}
}
