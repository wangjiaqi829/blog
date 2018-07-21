package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

// 标签
type Tag struct {
	Id         uint      `json:"id"`         // 标签Id
	Name       string    `json:"name"`       // 标签名称
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateTime time.Time `json:"updateTime"` // 更新时间
}



func (m *Tag) TableName() string {
	return TableName("tag")
}

func (m *Tag) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
