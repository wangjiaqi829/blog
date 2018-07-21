package models

import "time"

type Series struct {
	Id         uint      `orm:"pk" json:"id"`         //	系列ID
	Name       string    `json:"name"`       // 系列名称
	Sequence   int       `json:"sequence"`   // 排序 sequence越大越靠前
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateTime time.Time `json:"updateTime"` //	更新时间
}


