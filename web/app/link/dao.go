package link

import (
	"github.com/Seven4X/link/web/library/store/db"
	"github.com/xormplus/xorm"
)

type Dao struct {
	*xorm.Engine
}

var (
	ListLinkWithUser = `select "link.id", "link.link", "link.title", "link.l_group", "link.tags", "link.agree", "link.disagree", "link.create_time",
	b.is_like
	from link left join user_vote b on link.id = b.id and b.type='l' and b.user_id= ? 
	left join (select link_id,content,id from t_comment where t_comment.link_id = link.id order by agree desc limit 1) t on t.link_id = link.id  
	where a.id = ? limit ?`
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

func (dao *Dao) ListLink(req *ListLinkRequest) (links []Link, total int64, err error) {
	total, err = dao.countLink(req)
	if total < 1 {
		return nil, 0, err
	}
	links = make([]Link, 0)
	err = dao.Table("link").
		Cols(append([]string{"t_user.id", "t_user.name", "t_user.avatar"}, BaseColumn...)...).
		Join("left", "t_user", "t_user.id=link.create_by").
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

//
func (dao *Dao) ListLinkJoinHotComment(req *ListLinkRequest) (links []Link, total int64, err error) {
	total, err = dao.countLink(req)
	if total < 1 {
		return nil, 0, err
	}
	err = dao.Table("link").
		Cols(append([]string{"t_user.id", "t_user.name", "t_user.avatar"}, BaseColumn...)...).
		Join("left", "t_user", "t_user.id=link.create_by").
		Join("left", "", "").
		Find(&links)
	return links, total, err
}
func (dao *Dao) ListLinkJoinCommentAndUserVote(req *ListLinkRequest) (links []Link, total int64, err error) {
	total, err = dao.countLink(req)
	if total < 1 {
		return nil, 0, err
	}
	err = dao.SQL(ListLinkWithUser, req.UserId, req.Tid, req.Size).Find(&links)
	return links, total, err
}
