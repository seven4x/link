package topic

import (
	"time"
)

type Topic struct {
	Id   int    `json:"id" xorm:"pk autoincr"`
	Name string `json:"name"`
	Tags string `json:"tags"`
	Icon string `json:"icon"`
	//1 公开编辑， 2个人 3 团队
	Scope     int       `json:"scope"`
	Lang      string    `json:"lang"`
	CreateBy  int       `json:"createBy"`
	Score     int       `json:"score"`
	Agree     int       `json:"agree"`
	Disagree  int       `json:"disagree"`
	ShortCode string    `json:"shortCode"`
	CreatedAt time.Time `json:"createdAt" xorm:"create_time created"`
	UpdateAt  time.Time `json:"updateAt" xorm:"update_time updated"`
	DeletedAt time.Time `json:"-" xorm:"delete_time deleted"`
}

type TopicRel struct {
	Aid        int       `json:"aid"`      // 上 左
	Bid        int       `json:"bid"`      // 下 右
	Position   int       `json:"position"` //1 上下 2 左右
	CreateBy   int       `json:"createBy"`
	Predicate  string    `json:"predicate"`
	CreateTime time.Time `json:"createTime" xorm:"create_time created"`
	DeleteTime time.Time `json:"deleteTime" xorm:"update_time deleted"`
}

type TopicAlias struct {
	Alias   string
	TopicId int
}

type (
	HotTopic struct {
		Id         int
		Expire     time.Time
		CreateTime time.Time
		UpdateTime time.Time
		DeleteTime time.Time
	}
)
