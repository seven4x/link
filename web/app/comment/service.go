package comment

import (
	"github.com/Seven4X/link/web/app/messages"
	"github.com/Seven4X/link/web/app/risk"
	"github.com/Seven4X/link/web/app/util"
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

func (s *Service) SaveNewComment(req *NewCommentRequest) (id int, errs *util.Err) {
	comment := &Comment{
		LinkId:   req.LinkId,
		Content:  risk.SafeUserText(req.Content),
		CreateBy: req.CreateBy,
	}

	if _, err := s.dao.InsertOne(comment); err != nil {
		return -1, util.NewError(messages.GlobalErrorAboutDatabase)
	}
	s.dao.GrowCommentCnt(req.LinkId)
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
	//暂不事物控制
	affected, err := s.dao.Delete(comment)
	s.dao.DisGrowCommentCnt(linkId)
	return affected, err
}
