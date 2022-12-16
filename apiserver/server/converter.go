package server

import (
	"github.com/seven4x/link/api"
	"github.com/seven4x/link/db"
	"strconv"
)

func ConvertRequestToTopicModel(req *api.CreateTopicRequest) (topic *db.Topic, rel *db.TopicRel) {
	topic = &db.Topic{}
	topic.Name = req.Name
	topic.Tags = req.Tags
	topic.Lang = req.Lang
	topic.Scope = req.Scope

	rel = &db.TopicRel{}
	rel.Aid = req.RefTopicId
	rel.Position = req.Position
	rel.Predicate = req.Predicate
	return topic, rel
}
func ToLink(req *api.CreateLinkRequest) (link *db.Link) {
	topicId, _ := strconv.Atoi(req.TopicId)
	link = &db.Link{
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
