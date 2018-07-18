package models

import "time"

// 主题
type Topic struct {
	Id         uint      `json:"id"`         //	主题ID
	Name       string    `json:"name"`       // 主题名称
	Sequence   int       `json:"sequence"`   // 排序 sequence越大越靠前
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateTime time.Time `json:"updateTime"` //	更新时间
}
