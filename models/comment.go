package models

import "time"

// Comment 评论, 也称回复
type Comment struct {
	ID          uint
	Content     string     `json:"content"`
	HTMLContent string     `json:"htmlContent"`
	ContentType int        `json:"contentType"`
	ParentID    uint       `json:"parentID"`   //直接父评论的ID
	Parents     []Comment  `json:"parents"`    //所有的父评论
	SourceID    uint       `json:"sourceID"`   //话题或投票的ID
	Status      int        `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
