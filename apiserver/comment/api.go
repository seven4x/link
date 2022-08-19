package comment

import (
	"github.com/seven4x/link/web/user"
)

type NewCommentRequest struct {
	Content string `validate:"max=140"`
	LinkId  int
	//
	CreateBy int
}

type ListCommentRequest struct {
	Prev   int
	SortBy string
	Size   int
	UserId int
	LinkId int
}

type ListCommentResponse struct {
	Id       int    `json:"id" `
	LinkId   int    `json:"linkId"`
	Content  string `json:"content" `
	Score    int    `json:"score"  `
	Agree    int    `json:"agree"  `
	Disagree int    `json:"disagree"  `
	//之所以返回给前端Unix时间，时间显示样式由前端决定，为：1秒前，2天前，人类容易阅读的文本
	CreateTime int          `json:"createTime"  `
	Creator    user.Account `json:"creator" `
	IsLike     rune         `json:"isLike"`
}

func BuildListCommentFromModel(model *CommentUser) (res *ListCommentResponse) {
	model.Creator.Id = model.CreateBy
	return &ListCommentResponse{
		Id:         model.Id,
		LinkId:     model.LinkId,
		Content:    model.Content,
		Score:      model.Score,
		Agree:      model.Agree,
		Disagree:   model.Disagree,
		CreateTime: int(model.CreateTime.Unix()),
		Creator:    model.Creator,
		IsLike:     model.IsLike,
	}

}
