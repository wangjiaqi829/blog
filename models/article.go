package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

// 文章
type Article struct {
	Id           uint      `orm:"pk"  json:"id"`			 // 文章Id
	Title        string    `json:"name"`		 // 文章标题
	Content      string    `json:"content"`      // 原始内容
	HTMLContent  string    `json:"htmlContent"`  // Html文件
	Status       int       `json:"status"`       // 状态
	BrowseCount  uint      `json:"browseCount"`  // 浏览次数
	CommentCount uint      `json:"commentCount"` // 评论次数
	CreateTime   time.Time `json:"createTime"`   // 创建时间
	UpdateTime   time.Time `json:"updateTime"`   // 更新时间
}


func (m *Article) TableName() string {
	return TableName("article")
}

func (m *Article) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Article) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Article) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Article) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
