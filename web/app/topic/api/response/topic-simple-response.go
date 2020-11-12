package response

import "github.com/Seven4X/link/web/app/topic/model"

type TopicSimple struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func TopicSimpleOfModel(topic *model.Topic) (res *TopicSimple) {
	if topic == nil {
		return nil
	}
	res = &TopicSimple{
		Name: topic.Name,
		Id:   topic.Id,
	}

	return
}

func ModelToTopicSimple(topics []model.Topic) (res []*TopicSimple) {
	topic := make([]*TopicSimple, len(topics))
	for i, v := range topics {
		topic[i] = TopicSimpleOfModel(&v)
	}
	return topic
}
