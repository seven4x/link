package comment

import (
	"github.com/Seven4X/link/web/app/user"
	"time"
)

type Comment struct {
	Id         int       `json:"id" xorm:"pk autoincr"`
	LinkId     int       `json:"link_id"`
	Context    string    `json:"context" xorm:"varchar(240)"`
	Score      int       `json:"score" xorm:"default 0"`
	Agree      int       `json:"agree"  xorm:"default 0"`
	Disagree   int       `json:"disagree"  xorm:"default 0"`
	CreateBy   int       `json:"create_by"`
	CreateTime time.Time `json:"create_time"  xorm:"create_time created"`
	UpdateAt   time.Time `json:"update_time" xorm:"update_time updated"`
	DeletedAt  time.Time `json:"delete_time" xorm:"delete_time deleted"`
	IsLike     rune      `xorm:"<-"`
}

type CommentUser struct {
	Comment
	Creator user.Account `xorm:"extends"`
}
