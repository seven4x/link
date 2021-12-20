package link

import (
	"github.com/seven4x/link/web/store"
	"strconv"
	"xorm.io/xorm"
)

type Dao struct {
	*xorm.Engine
}

var (
	BaseColumn = []string{"link.id", "link.link", "link.title", "link.l_group", "link.tags", "link.agree", "link.disagree", "link.comment_cnt", "link.create_time"}
)

func NewDao() (dao *Dao) {
	dao = &Dao{store.NewDb()}
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
	start := (req.Page - 1) * req.Size
	if start < 0 {
		start = 0
	}
	sess := dao.Table("link").
		Cols(append([]string{"account.nick_name", "account.avatar"}, BaseColumn...)...).
		Join("left", "account", "account.id=link.create_by").
		Where("link.topic_id=?", req.Tid).
		And("link.id>?", req.Prev)
	if req.FilterMy {
		sess.And("link.create_by=?", req.UserId)
	}
	if req.OrderBy != "" {
		sess.OrderBy(req.OrderBy)
	}
	err = sess.
		Limit(req.Size, start).
		Find(&links)
	return links, total, err
}

func (dao *Dao) ListLinkJoinUserVote(req *ListLinkRequest) (links []LinkUser, total int64, err error) {
	total, err = dao.countLink(req)
	if total < 1 {
		return nil, 0, err
	}
	err = dao.Table("link").
		Cols(append([]string{"account.id", "account.nick_name", "account.avatar", "user_vote.is_like"}, BaseColumn...)...).
		Join("left", "account", "account.id=link.create_by").
		Join("left", "user_vote", strconv.Itoa(req.UserId)+"=user_vote.user_id and user_vote.type='l' and user_vote.id=link.id").
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
