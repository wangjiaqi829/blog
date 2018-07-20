package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

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
	orm.RegisterModel(new(Article),new(models.Comment),new(models.Tag),new(models.User))

	// 开启 ORM 调试模式
	orm.Debug = true

	// 自动建表
	orm.RunSyncdb("default", false, true)


	if beego.AppConfig.String("RunMode") == "dev" {
		orm.Debug = false
	}
}
