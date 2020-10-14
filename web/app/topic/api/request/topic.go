package request

import "github.com/Seven4X/link/web/app/topic/model"

type NewTopicReq struct {
	Name       string `json:"name" validate:"required,min=2,max=140"`
	RefTopicId int    `json:"refTopicId" validate:"required"`
	Position   int    `json:"position" validate:"oneof=1 2 3 4"`
	Predicate  string `json:"predicate"`
	Tags       string `json:"tags"`
}

func (req *NewTopicReq) ToTopic() (topic *model.Topic) {
	topic = &model.Topic{}
	topic.Name = req.Name
	topic.Tags = req.Tags
	return topic
}
