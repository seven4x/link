package db

import (
	"github.com/seven4x/link/api"
	"strconv"
	"time"
)

type Link struct {
	Id           int       `json:"id" xorm:"pk autoincr"`
	Link         string    `json:"link"`
	Title        string    `json:"title"`
	Group        string    `json:"group" xorm:"l_group"`
	Tags         string    `json:"tags" xorm:"varchar(140)"`
	From         int       `json:"from" xorm:"l_from char(1)"`
	TopicId      int       `json:"topicId"`
	Score        int       `json:"score"`
	Agree        int       `json:"agree"`
	Disagree     int       `json:"disagree"`
	CommentCnt   int       `json:"commentCnt "`
	CreateAt     time.Time `json:"create_time" xorm:"create_time created"`
	UpdateAt     time.Time `json:"update_time" xorm:"update_time updated"`
	DeletedAt    time.Time `json:"delete_time" xorm:"delete_time deleted"`
	CreateBy     int       `json:"create_by"`
	FirstComment string    `xorm:"-"`
	IsLike       int       `xorm:"<-"`
}

type WithUser struct {
	Link    `xorm:"extends"`
	Creator Account `xorm:"extends "`
}

var (
	BaseColumn = []string{"link.id", "link.link", "link.title", "link.l_group", "link.tags", "link.agree", "link.disagree", "link.comment_cnt", "link.create_time"}
)

func (dao *Dao) SaveLink(link *Link) (i int, err error) {
	_, err = dao.Insert(link)
	return link.Id, err
}

func (dao *Dao) ListLink(req *api.ListLinkRequest) (links []WithUser, err error) {

	links = make([]WithUser, 0)

	sess := dao.Table("link").
		Cols(append([]string{"account.nick_name", "account.avatar"}, BaseColumn...)...).
		Join("left", "account", "account.id=link.create_by").
		Where("link.topic_id=?", req.Tid)
	if req.Prev != 0 {
		sess.And("link.id < ?", req.Prev)
	}
	if req.FilterMy {
		sess.And("link.create_by=?", req.UserId)
	}
	if req.OrderBy != "" {
		sess.OrderBy(req.OrderBy)
	}
	err = sess.Limit(req.Size).Find(&links)
	return links, err
}

func (dao *Dao) ListLinkJoinUserVote(req *api.ListLinkRequest) (links []WithUser, total int64, err error) {
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

func (dao *Dao) countLink(req *api.ListLinkRequest) (total int64, err error) {
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

func (dao *Dao) GetRecentLinks(prev int) ([]Link, error) {
	res := make([]Link, 0)
	sql := "SELECT ID,link,title,tags,create_time FROM LINK WHERE id < ? ORDER BY id DESC  limit 20"
	if prev == 0 {
		sql = "SELECT ID,link,title,tags,create_time FROM LINK   ORDER BY id DESC  limit 20"
	}
	err := dao.SQL(sql, prev).Find(&res)
	return res, err
}
