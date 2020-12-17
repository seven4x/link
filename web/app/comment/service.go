package comment

type Service struct {
	dao *Dao
}

func NewService() (service *Service) {
	service = &Service{
		dao: NewDao(),
	}
	return
}

func (service *Service) Save(comment *Comment) (id int, err error) {
	_, err = service.dao.InsertOne(comment)
	return comment.Id, err
}

func (service *Service) ListHotCommentByLinkId(ids []interface{}) ([]Comment, error) {

	return service.dao.ListHotCommentByLinkId(ids)
}
