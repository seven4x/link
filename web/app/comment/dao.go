package comment

import (
	"github.com/Seven4X/link/web/library/store/db"
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

func (dao *Dao) ListHotCommentByLinkId(ids []interface{}) ([]Comment, error) {
	res := make([]Comment, 0)
	err := dao.Table("t_comment").
		Cols("t_comment.context", "t_comment.create_time", "t_comment.link_id",
			"t_user.id", "t_user.avatar", "agree", "t_user.name").
		Join("left", "t_user", "t_user.id=t_comment.create_by").
		Where(builder.In("t_comment.link_id", ids...)).
		OrderBy("agree desc").Limit(1, 0).
		Find(&res)
	return res, err
}

func (dao *Dao) ListComment(req *ListCommentRequest) (res []*Comment, hasMore bool, err error) {
	res = make([]*Comment, 0)
	cols := []string{"t_comment.*", "t_user.id,t_user.name,t_user.avatar"}
	if req.UserId != 0 {
		cols = append(cols, "user_vote.is_like")
	}
	session := dao.Table("t_comment").
		Cols(cols...).
		Join("left", "t_user", "t_user.id = t_comment.create_by")
	if req.UserId != 0 {
		session.Join("left", "user_vote", strconv.Itoa(req.UserId)+"=user_vote.user_id and user_vote.type='c' and user_vote.id=t_comment.id")
	}
	session.Where("t_comment.link_id=?", req.LinkId).And("t_comment.id>?", req.Prev)
	if req.SortBy == "newest" {
		session.OrderBy("t_comment.create_time desc")
	} else if req.SortBy == "hot" {
		session.OrderBy("t_comment.agree desc ")
	}
	session.Limit(req.Size+1, 0).Find(&res)
	return res, len(res) == req.Size, err
}
