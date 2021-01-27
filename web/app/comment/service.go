package comment

import (
	"github.com/Seven4X/link/web/app/risk"
	"github.com/Seven4X/link/web/lib/api"
	"github.com/Seven4X/link/web/lib/api/messages"
)

type Service struct {
	dao *Dao
}

func NewService() (service *Service) {
	service = &Service{
		dao: NewDao(),
	}
	return
}
func (s *Service) Save(comment *Comment) (id int, err error) {
	_, err = s.dao.InsertOne(comment)
	return comment.Id, err
}

func (s *Service) SaveNewComment(req *NewCommentRequest) (id int, errs *api.Err) {
	comment := &Comment{
		LinkId:   req.LinkId,
		Context:  risk.SafeUserText(req.Content),
		CreateBy: req.CreateBy,
	}

	if _, err := s.dao.InsertOne(comment); err != nil {
		return -1, api.NewError(messages.GlobalErrorAboutDatabase)
	}
	return comment.Id, nil
}
func (s *Service) ListHotCommentByLinkId(ids []interface{}) ([]CommentUser, error) {

	return s.dao.ListHotCommentByLinkId(ids)
}

func (s *Service) ListComment(req *ListCommentRequest) (res []*ListCommentResponse, hasMore bool, err error) {
	models, hasMore, err := s.dao.ListComment(req)
	res = make([]*ListCommentResponse, 0)
	for _, model := range models {
		res = append(res, BuildListCommentFromModel(model))
	}
	return res, hasMore, err

}

func (s *Service) DeleteComment(linkId int, commentId int) (int64, error) {
	comment := &Comment{
		Id:     commentId,
		LinkId: linkId,
	}
	affected, err := s.dao.Delete(comment)
	return affected, err
}
