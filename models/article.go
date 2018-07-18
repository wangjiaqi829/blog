package models

import "time"

// 文章
type Article struct {
	Id           uint      `json:"id"`			 // 文章Id
	Title        string    `json:"name"`		 // 文章标题
	Content      string    `json:"content"`      // 原始内容
	HTMLContent  string    `json:"htmlContent"`  // Html文件
	Status       int       `json:"status"`       // 状态
	BrowseCount  uint      `json:"browseCount"`  // 浏览次数
	CommentCount uint      `json:"commentCount"` // 评论次数
	CreateTime   time.Time `json:"createTime"`   // 创建时间
	UpdateTime   time.Time `json:"updateTime"`   // 更新时间
}

// 文章状态
const (
	h = iota
)
