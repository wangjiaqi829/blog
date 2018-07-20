package models

import (
<<<<<<< HEAD
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {

	err := beego.LoadAppConfig("ini", "./src/ty/conf/app.conf")
	if err != nil {
		beego.Error("loadAppConfig error:", err.Error())
	}
=======
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"strings"
	"time"
)

func init() {
>>>>>>> 07442226966015f0ad8778f6ebd5d3ffdd491a79
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
<<<<<<< HEAD

	beego.Info(dbhost, dbuser)
	dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Local"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dburl)

	//将定义的iot Model 进行注册
	orm.RegisterModel(new(Article),new(models.Comment),new(models.Tag),new(models.User))
=======
	dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"

	//注册数据库类型
	orm.RegisterDriver("mysql", orm.DRMySQL)

	//设置数据库时区
	orm.DefaultTimeLoc = time.UTC
	//ORM 必须注册一个别名为 default 的数据库，作为默认使用
	orm.RegisterDataBase("default", "mysql", dburl)

	//将定义的iot Model 进行注册
	orm.RegisterModel(new(Article),new(Comment),new(Tag),new(User))
>>>>>>> 07442226966015f0ad8778f6ebd5d3ffdd491a79

	// 开启 ORM 调试模式
	orm.Debug = true

	// 自动建表
	orm.RunSyncdb("default", false, true)

<<<<<<< HEAD

	if beego.AppConfig.String("RunMode") == "dev" {
		orm.Debug = false
	}
}
=======
	orm.RegisterModel()
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

//返回带前缀的表名
func TableName(str string) string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("dbprefix"), str)
}

>>>>>>> 07442226966015f0ad8778f6ebd5d3ffdd491a79
