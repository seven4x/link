package request

import "github.com/Seven4X/link/web/app/topic/model"

type CreateTopicRequest struct {
	Name       string `json:"name" validate:"required,min=2,max=140"`
	RefTopicId int    `json:"refId" validate:"required"`
	Position   int    `json:"position" validate:"oneof=1 2 3 4"`
	Predicate  string `json:"refDesc"`
	Tags       string `json:"tags"`
	Lang       string `json:"lang"`
}

func (req *CreateTopicRequest) ToTopic() (topic *model.Topic, rel *model.TopicRel) {
	topic = &model.Topic{}
	topic.Name = req.Name
	topic.Tags = req.Tags
	topic.Lang = req.Lang

	rel = &model.TopicRel{}
	rel.Aid = req.RefTopicId
	rel.Position = req.Position
	rel.Predicate = req.Predicate
	return topic, rel
}
