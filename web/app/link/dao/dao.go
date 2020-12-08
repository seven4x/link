package dao

import (
	"github.com/Seven4X/link/web/app/link/model"
	"github.com/Seven4X/link/web/app/link/server/request"
	"github.com/Seven4X/link/web/library/store/db"
	"github.com/xormplus/xorm"
)

type Dao struct {
	*xorm.Engine
}

var (
	ListLinkWithUser = `select a.*,b.is_like
	from link a left join user_vote b on a.id = b.id and b.type='l' and b.user_id=? 
	where a.id = ? limit ?`
)

func New() (dao *Dao) {
	dao = &Dao{db.NewDb()}
	return
}

func (dao *Dao) Save(link *model.Link) (i int, err error) {
	_, err = dao.Insert(link)
	return link.Id, err
}

func (dao *Dao) ListLinkWithoutUser(req *request.ListLinkRequest) (links []model.Link, total int64, err error) {

	link := model.Link{}
	link.TopicId = req.Tid
	total, err = dao.Where("id>?", req.Prev).Count(link)
	if err != nil {
		return nil, 0, err
	}
	err = dao.Where("topic_id=?", req.Tid).And("id>?", req.Prev).Limit(req.Size, 0).Find(&links)
	return links, total, err
}
func (dao *Dao) ListLinkWithUser(req *request.ListLinkRequest) (links []model.Link, total int64, err error) {
	link := model.Link{}
	link.TopicId = req.Tid
	total, err = dao.Where("id>?", req.Prev).Count(link)
	if err != nil {
		return nil, 0, err
	}
	err = dao.SQL(ListLinkWithUser, req.UserId, req.Tid, req.Size).Find(&links)
	return links, total, err
}
