package models

import "time"

// 标签
type Tag struct {
	Id         uint      `json:"id"`         // 标签Id
	Name       string    `json:"name"`       // 标签名称
	CreateTime time.Time `json:"createTime"` // 创建时间
	UpdateTime time.Time `json:"updateTime"` // 更新时间
}
