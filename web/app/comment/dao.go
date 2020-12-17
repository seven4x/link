package comment

import (
	"github.com/Seven4X/link/web/library/store/db"
	"github.com/xormplus/builder"
	"github.com/xormplus/xorm"
)

type Dao struct {
	*xorm.Engine
}

const (
	ListHotCommentByLinkIds = `select  
		a.context,a.create_time,
		b.avatar 
		from t_comment a 
		left join t_user b on a.create_by = b.id
		where a.id in (?)
`
)

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
