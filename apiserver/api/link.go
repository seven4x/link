package api

import (
	"time"
)

type ListLinkResponse struct {
	Id         int    `json:"id"`
	Link       string `json:"link"`
	Title      string `json:"title"`
	Agree      int    `json:"agree"`
	Disagree   int    `json:"disagree"`
	Group      string `json:"group"`
	Score      int
	CreateTime time.Time   `json:"createTime"`
	IsLike     int         `json:"isLike"`
	HotComment *HotComment `json:"hotComment"`
	CreateBy   *Creator
	//评论数
	CommentCount int `json:"commentCount"`
	ClickCount   int `json:"clickCount"`
}

type HotComment struct {
	UserId  int    `json:"userId"`
	Content string `json:"content"`
	Avatar  string `json:"avatar"`
}
type Creator struct {
	Name   string
	UserId int
	Avatar string
}

type ListLinkRequest struct {
	Prev     int    `query:"prev"`
	Size     int    `query:"size"`
	Tid      int    `validate:"required" query:"tid"`
	Group    string `query:"group"'`
	OrderBy  string
	UserId   int
	FilterMy bool
}

type CreateLinkRequest struct {
	TopicId string `validate:"required"`
	Title   string `validate:"required"`
	Url     string `validate:"required"`
	Comment string `validate:"max=140"`
	Group   string `validate:"max=15"`
	Tags    string `validate:"max=140"`
	//1 web 2 chrome
	From int `validate:"oneof=1 2"`
}

type RecentVO struct {
	Title      string
	Id         int
	Url        string
	Tags       string
	CreateTime time.Time
}
