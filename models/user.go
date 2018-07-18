package models

// 用户
type User struct {
	Id     uint   `json:"id"`     // 用户Id
	Name   string `json:"name"`   // 用户昵称
	Sign   string `json:"sign"`   // 用户签名
	Avatar string `json:"avatar"` // 用户头像
	GitHub string `json:"gitHub"` // github地址
}
