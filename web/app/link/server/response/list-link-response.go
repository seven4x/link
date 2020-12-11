package response

import "time"

type ListLinkResponse struct {
	Id         int    `json:"id"`
	Link       string `json:"link"`
	Title      string `json:"title"`
	Agree      int    `json:"agree"`
	Disagree   int    `json:"disagree"`
	Group      string `json:"group"`
	Score      int
	CreateTime time.Time `json:"createTime"`
	IsLike     rune      `json:"isLike"`
	HotComment struct {
		UserId  int
		Context string
		Avatar  string
	}
	CreateBy struct {
		Name   string
		UserId int
		Avatar string
	}
	//评论数
	CommentCount int `json:"commentCount"`
	ClickCount   int `json:"clickCount"`
}
