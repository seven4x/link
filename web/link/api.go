package link

import (
	"strconv"
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

func BuildLinkResponseOfModel(m *WithUser) (res *ListLinkResponse) {

	res = &ListLinkResponse{
		Id:           m.Id,
		Link:         m.Link.Link,
		Title:        m.Title,
		Agree:        m.Agree,
		Disagree:     m.Disagree,
		Group:        m.Group,
		Score:        m.Score,
		CreateTime:   m.CreateAt,
		IsLike:       m.IsLike,
		HotComment:   nil,
		CreateBy:     nil,
		CommentCount: m.CommentCnt,
		ClickCount:   0,
	}
	res.CreateBy = &Creator{
		Name:   m.Creator.NickName,
		UserId: m.Creator.Id,
		Avatar: m.Creator.Avatar,
	}
	return res
}

type ListLinkRequest struct {
	Prev     int `query:"prev"`
	Page     int
	Size     int
	Tid      int `validate:"required" query:"tid"`
	Group    string
	UserId   int
	OrderBy  string
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

func (req *CreateLinkRequest) ToLink() (link *Link) {
	topicId, _ := strconv.Atoi(req.TopicId)
	link = &Link{
		Id:           0,
		Link:         req.Url,
		Title:        req.Title,
		TopicId:      topicId,
		Score:        0,
		Agree:        0,
		Disagree:     0,
		From:         req.From,
		Group:        req.Group,
		Tags:         req.Tags,
		FirstComment: req.Comment,
		CommentCnt:   1,
	}
	return
}

type RecentVO struct {
	Title      string
	Id         int
	Url        string
	Tags       string
	CreateTime time.Time
}
