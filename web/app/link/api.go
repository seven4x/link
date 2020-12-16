package link

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
	HotComment *HotComment
	CreateBy   *Creator
	//评论数
	CommentCount int `json:"commentCount"`
	ClickCount   int `json:"clickCount"`
}

type HotComment struct {
	UserId  int
	Context string
	Avatar  string
}
type Creator struct {
	Name   string
	UserId int
	Avatar string
}

func BuildLinkResponseOfModel(m *Link) (res *ListLinkResponse) {

	res = &ListLinkResponse{
		Id:           m.Id,
		Link:         m.Link,
		Title:        m.Title,
		Agree:        m.Agree,
		Disagree:     m.Disagree,
		Group:        m.Group,
		Score:        m.Score,
		CreateTime:   m.CreateAt,
		IsLike:       m.IsLike,
		HotComment:   nil,
		CreateBy:     nil,
		CommentCount: 0,
		ClickCount:   0,
	}
	res.CreateBy = &Creator{
		Name:   m.Creator.Name,
		UserId: m.Creator.Id,
		Avatar: m.Creator.Avatar,
	}
	return res
}

type NewCommentRequest struct {
	Content string `validate:"max=140"`
	LinkId  int
	//
	CreateBy int
}

type ListLinkRequest struct {
	Prev   int
	Size   int
	Tid    int `validate:"required"`
	Group  string
	UserId int
}

type CreateLinkRequest struct {
	TopicId int    `validate:"required"`
	Title   string `validate:"required"`
	Url     string `validate:"required"`
	Comment string `validate:"max=140"`
	Group   string `validate:"max=15"`
	Tags    string `validate:"max=140"`
	//1 web 2 chrome
	From int `validate:"oneof=1 2"`
}

func (req *CreateLinkRequest) ToLink() (link *Link) {
	link = &Link{
		Id:           0,
		Link:         req.Url,
		Title:        req.Title,
		TopicId:      req.TopicId,
		Score:        0,
		Agree:        0,
		Disagree:     0,
		From:         req.From,
		Group:        req.Group,
		Tags:         req.Tags,
		FirstComment: req.Comment,
	}
	return
}
