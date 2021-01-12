package link

import (
	"github.com/Seven4X/link/web/library/store/db"
	"github.com/xormplus/xorm"
	"strconv"
)

type Dao struct {
	*xorm.Engine
}

var (
	BaseColumn = []string{"link.id", "link.link", "link.title", "link.l_group", "link.tags", "link.agree", "link.disagree", "link.create_time"}
)

func NewDao() (dao *Dao) {
	dao = &Dao{db.NewDb()}
	return
}

func (dao *Dao) Save(link *Link) (i int, err error) {
	_, err = dao.Insert(link)
	return link.Id, err
}

func (dao *Dao) ListLink(req *ListLinkRequest) (links []LinkUser, total int64, err error) {
	total, err = dao.countLink(req)
	if total < 1 {
		return nil, 0, err
	}
	links = make([]LinkUser, 0)
	err = dao.Table("link").
		Cols(append([]string{"account.id", "account.name", "account.avatar"}, BaseColumn...)...).
		Join("left", "account", "account.id=link.create_by").
		Where("link.topic_id=?", req.Tid).
		And("link.id>?", req.Prev).
		Limit(req.Size, 0).
		Find(&links)
	return links, total, err
}
func (dao *Dao) countLink(req *ListLinkRequest) (total int64, err error) {
	link := Link{}
	link.TopicId = req.Tid
	total, err = dao.Table("link").
		Where("id>?", req.Prev).Count(&link)
	if err != nil {
		return 0, err
	} else {
		return total, err
	}

}

func (dao *Dao) ListLinkJoinUserVote(req *ListLinkRequest) (links []LinkUser, total int64, err error) {
	total, err = dao.countLink(req)
	if total < 1 {
		return nil, 0, err
	}
	err = dao.Table("link").
		Cols(append([]string{"account.id", "account.name", "account.avatar", "user_vote.is_like"}, BaseColumn...)...).
		Join("left", "account", "account.id=link.create_by").
		Join("left", "user_vote", strconv.Itoa(req.UserId)+"=user_vote.user_id and user_vote.type='l' and user_vote.id=link.id").
		Where("link.topic_id=?", req.Tid).
		And("link.id>?", req.Prev).
		Limit(req.Size, 0).
		Find(&links)
	return links, total, err
}
