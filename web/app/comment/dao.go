package comment

import (
	"github.com/Seven4X/link/web/library/store/db"
	"github.com/xormplus/xorm"
)

type Dao struct {
	*xorm.Engine
}

const (
	ListHotCommentByLinkIds = `select  
		a.context,a.create_time,
		b.avatar as "creator.avatar",b.avatar as creator_avatar
		from t_comment a 
		left join t_user b on a.create_by = b.id
		where a.link_id in (?)
`
)

func NewDao() (dao *Dao) {
	dao = &Dao{db.NewDb()}
	return
}

func (dao *Dao) ListHotCommentByLinkId(ids []int) ([]Comment, error) {
	res := make([]Comment, 0)
	err := dao.SQL(ListHotCommentByLinkIds, 1).Find(&res)
	return res, err
}
