package service

import (
	"github.com/seven4x/link/api"
	"github.com/seven4x/link/db"
	"strconv"
)

func BuildListCommentFromModel(model *db.CommentUser) (res *api.ListCommentResponse) {
	model.Creator.Id = model.CreateBy
	return &api.ListCommentResponse{
		Id:         model.Id,
		LinkId:     model.LinkId,
		Content:    model.Content,
		Score:      model.Score,
		Agree:      model.Agree,
		Disagree:   model.Disagree,
		CreateTime: int(model.CreateTime.Unix()),
		Creator: api.AccountInfo{
			Id:       model.Creator.Id,
			Name:     model.Creator.UserName,
			NickName: model.Creator.NickName,
			Avatar:   model.Creator.Avatar,
		},
		IsLike: model.IsLike,
	}

}
func BuildLinkResponseOfModel(m *db.WithUser) (res *api.ListLinkResponse) {

	res = &api.ListLinkResponse{
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
	res.CreateBy = &api.Creator{
		Name:   m.Creator.NickName,
		UserId: m.Creator.Id,
		Avatar: m.Creator.Avatar,
	}
	return res
}

func BuildDetailFromModel(topic *db.Topic) (res *api.TopicDetail) {
	if topic == nil {
		return nil
	}
	res = &api.TopicDetail{
		Name:      topic.Name,
		Id:        topic.Id,
		Icon:      topic.Icon,
		ShortCode: topic.ShortCode,
	}

	return
}

func buildSimpleFromModel(topic *db.Topic) (res *api.SnapShot) {
	if topic == nil {
		return nil
	}
	res = &api.SnapShot{
		Name:      topic.Name,
		Id:        strconv.Itoa(topic.Id),
		ShortCode: topic.ShortCode,
	}

	return
}

func ConvertModelToTopicSimple(topics []db.Topic) (res []*api.SnapShot) {
	topic := make([]*api.SnapShot, len(topics))
	for i, v := range topics {
		topic[i] = buildSimpleFromModel(&v)
	}
	return topic
}
