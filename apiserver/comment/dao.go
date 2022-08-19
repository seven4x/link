package comment

import (
	"fmt"
	"github.com/seven4x/link/web/store"
	"github.com/seven4x/link/web/util"
	"strconv"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type Dao struct {
	*xorm.Engine
}

func NewDao() (dao *Dao) {
	dao = &Dao{store.NewDb()}
	return
}

const listHostCommentSql = `select c0.*,a.avatar,a.nick_name
from comment c0
left join account a on c0.create_by = a.id
where c0.id in (
    select max(c1.id) id
    from comment c1
             inner join (select c2.link_id, max(c2.agree) agree
                         from comment c2
                         where %s
                           and c2.delete_time is null
                           and c2.agree > 0
                         group by c2.link_id) tmp on tmp.link_id = c1.link_id and tmp.agree = c1.agree
    group by c1.link_id
)`

func (dao *Dao) ListHotCommentByLinkId(ids []interface{}) ([]CommentUser, error) {
	res := make([]CommentUser, 0)
	str, _, _ := builder.ToSQL(builder.In("c2.link_id", ids...))
	sql := fmt.Sprintf(listHostCommentSql, str)
	err := dao.SQL(sql, ids...).Find(&res)
	return res, err
}

func (dao *Dao) ListComment(req *ListCommentRequest) (res []*CommentUser, hasMore bool, err error) {
	res = make([]*CommentUser, 0)
	cols := []string{"comment.*", "account.nick_name,account.avatar,account.user_name"}
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
		util.Error(err.Error())
	}
}
func (dao *Dao) DisGrowCommentCnt(linkId int) {
	_, err := dao.Exec("update link set comment_cnt = comment_cnt - 1 where id=?  ", linkId)
	if err != nil {
		util.Error(err.Error())
	}
}
