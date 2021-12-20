package user

import "time"

type Account struct {
	Id       int    `json:"id" xorm:"pk autoincr"`
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Password string `json:"-"`
}

// 用户邀请码对照表
type RegisterCode struct {
	Id     int `xorm:"pk autoincr"`
	UserId int
	Code   string
}

//邀请码使用记录
type RegisterInfo struct {
	Id           int `xorm:"pk autoincr"`
	Code         string
	CreateBy     int
	UsedBy       int
	UsedUserName string
	UsedTime     time.Time
}

type UserVO struct {
	Account
}
