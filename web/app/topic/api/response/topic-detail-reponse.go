package response

import "github.com/Seven4X/link/web/app/topic/model"

type TopicDetail struct {
	Name       string `json:"name"`
	Id         int    `json:"id"`
	Icon       string `json:"icon"`
	CreateUser string `json:"createUser"`
}

func TopicDetailOfModel(topic *model.Topic) (res *TopicDetail) {
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
