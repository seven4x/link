package topic

type CreateTopicRequest struct {
	Name       string `json:"name" validate:"required,min=2,max=140"`
	RefTopicId int    `json:"refId" validate:"required"`
	Position   int    `json:"position" validate:"oneof=1 2 3 4"`
	Predicate  string `json:"refDesc"`
	Tags       string `json:"tags"`
	Lang       string `json:"lang"`
	Scope      int    `json:"scope" validate:"oneof=1 2 3"`
}

func (req *CreateTopicRequest) ToTopic() (topic *Topic, rel *TopicRel) {
	topic = &Topic{}
	topic.Name = req.Name
	topic.Tags = req.Tags
	topic.Lang = req.Lang
	topic.Scope = req.Scope

	rel = &TopicRel{}
	rel.Aid = req.RefTopicId
	rel.Position = req.Position
	rel.Predicate = req.Predicate
	return topic, rel
}

type TopicDetail struct {
	Name       string `json:"name"`
	Id         int    `json:"id"`
	Icon       string `json:"icon"`
	CreateUser string `json:"createUser"`
}

func TopicDetailOfModel(topic *Topic) (res *TopicDetail) {
	if topic == nil {
		return nil
	}
	res = &TopicDetail{
		Name: topic.Name,
		Id:   topic.Id,
		Icon: topic.Icon,
	}

	return
}

type TopicSimple struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func TopicSimpleOfModel(topic *Topic) (res *TopicSimple) {
	if topic == nil {
		return nil
	}
	res = &TopicSimple{
		Name: topic.Name,
		Id:   topic.Id,
	}

	return
}

func ModelToTopicSimple(topics []Topic) (res []*TopicSimple) {
	topic := make([]*TopicSimple, len(topics))
	for i, v := range topics {
		topic[i] = TopicSimpleOfModel(&v)
	}
	return topic
}
