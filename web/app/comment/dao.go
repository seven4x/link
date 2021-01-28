package comment

import (
	"github.com/Seven4X/link/web/lib/log"
	"github.com/Seven4X/link/web/lib/store/db"
	"github.com/xormplus/builder"
	"github.com/xormplus/xorm"
	"strconv"
)

type Dao struct {
	*xorm.Engine
}

func NewDao() (dao *Dao) {
	dao = &Dao{db.NewDb()}
	return
}

func (dao *Dao) ListHotCommentByLinkId(ids []interface{}) ([]CommentUser, error) {
	res := make([]CommentUser, 0)
	err := dao.Table("comment").
		Cols("comment.context", "comment.create_time", "comment.link_id",
			"account.avatar", "agree", "account.nick_name").
		Join("left", "account", "account.id=comment.create_by").
		Where(builder.In("comment.link_id", ids...)).
		OrderBy("agree desc").Limit(1, 0).
		Find(&res)
	return res, err
}

func (dao *Dao) ListComment(req *ListCommentRequest) (res []*CommentUser, hasMore bool, err error) {
	res = make([]*CommentUser, 0)
	cols := []string{"comment.*", "account.nick_name,account.avatar"}
	if req.UserId != 0 {
		cols = append(cols, "user_vote.is_like")
	}
	session := dao.Table("comment").
		Cols(cols...).
		Join("left", "account", "account.id = comment.create_by")
	if req.UserId != 0 {
		session.Join("left", "user_vote", strconv.Itoa(req.UserId)+"=user_vote.user_id and user_vote.type='c' and user_vote.id=comment.id")
	}
	session.Where("comment.link_id=?", req.LinkId).And("comment.id>?", req.Prev)
	if req.SortBy == "newest" {
		session.OrderBy("comment.create_time desc")
	} else if req.SortBy == "hot" {
		session.OrderBy("comment.agree desc ")
	}
	err = session.Limit(req.Size+1, 0).Find(&res)
	size := len(res)
	if size == 0 {
		return res, false, err
	}
	if size > req.Size {
		return res[:len(res)-1], len(res) > req.Size, err
	}
	return res, false, err
}

func (dao *Dao) GrowCommentCnt(linkId int) {
	_, err := dao.Exec("update link set comment_cnt = comment_cnt + 1 where id=? ", linkId)
	if err != nil {
		log.Error(err.Error())
	}
}
func (dao *Dao) DisGrowCommentCnt(linkId int) {
	_, err := dao.Exec("update link set comment_cnt = comment_cnt - 1 where id=?  ", linkId)
	if err != nil {
		log.Error(err.Error())
	}
}
