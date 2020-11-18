package request

import "github.com/Seven4X/link/web/app/link/model"

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

func (req *CreateLinkRequest) ToLink() (link *model.Link) {
	link = &model.Link{
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
