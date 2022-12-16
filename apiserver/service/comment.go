package service

import (
	"github.com/seven4x/link/api"
	"github.com/seven4x/link/app"
	"github.com/seven4x/link/db"
	"github.com/seven4x/link/service/risk"
)

func (s *Service) SaveComment(comment *db.Comment) (id int, err error) {
	_, err = s.Dao.InsertOne(comment)
	return comment.Id, err
}

func (s *Service) SaveNewComment(req *api.NewCommentRequest) (id int, errs *app.Err) {
	comment := &db.Comment{
		LinkId:   req.LinkId,
		Content:  risk.SafeUserText(req.Content),
		CreateBy: req.CreateBy,
	}

	if _, err := s.Dao.InsertOne(comment); err != nil {
		return -1, app.NewError(api.GlobalErrorAboutDatabase)
	}
	s.Dao.GrowCommentCnt(req.LinkId)
	return comment.Id, nil
}
func (s *Service) ListHotCommentByLinkId(ids []interface{}) ([]db.CommentUser, error) {

	return s.Dao.ListHotCommentByLinkId(ids)
}

func (s *Service) ListComment(req *api.ListCommentRequest) (res []*api.ListCommentResponse, hasMore bool, err error) {
	models, hasMore, err := s.Dao.ListComment(req)
	res = make([]*api.ListCommentResponse, 0)
	for _, model := range models {
		res = append(res, BuildListCommentFromModel(model))
	}
	return res, hasMore, err

}

func (s *Service) DeleteComment(linkId int, commentId int) (int64, error) {
	comment := &db.Comment{
		Id:     commentId,
		LinkId: linkId,
	}
	//暂不事物控制
	affected, err := s.Dao.Delete(comment)
	s.Dao.DisGrowCommentCnt(linkId)
	return affected, err
}
