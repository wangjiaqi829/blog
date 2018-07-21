package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Comment 评论, 也称回复
type Comment struct {
	ID          uint	  `orm:"pk" json:"id"`
	Content     string    `json:"content"`
	HTMLContent string    `json:"htmlContent"`
	ContentType int       `json:"contentType"`
	ParentID    uint      `json:"parentID"` //直接父评论的ID
	SourceID    uint      `json:"sourceID"` //话题或投票的ID
	Status      int       `json:"status"`
	CreateTime  time.Time `json:"created_at"`
	UpdateTime  time.Time `json:"updateTime"`
}

func (m *Comment) TableName() string {
	return TableName("comment")
}

func (m *Comment) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Comment) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Comment) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Comment) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Comment) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
