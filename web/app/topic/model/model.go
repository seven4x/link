package model

import (
	"time"
)

type Topic struct {
	Id        int       `json:"id" xorm:"pk autoincr"`
	Name      string    `json:"name"`
	Tags      string    `json:"tags"`
	Icon      string    `json:"icon"`
	Lang      string    `json:"lang"`
	CreateBy  int       `json:"createBy"`
	Score     int       `json:"score"`
	Agree     int       `json:"agree"`
	Disagree  int       `json:"disagree"`
	CreatedAt time.Time `json:"createdAt" xorm:"create_time created"`
	UpdateAt  time.Time `json:"updateAt" xorm:"update_time updated"`
	DeletedAt time.Time `json:"-" xorm:"delete_time deleted"`
}

type TopicRel struct {
	Aid        int       `json:"aid"`
	Bid        int       `json:"bid"`
	Position   int       `json:"position"` //1 上下 2 左右
	CreateBy   int       `json:"createBy"`
	Predicate  string    `json:"predicate"`
	CreateTime time.Time `json:"createTime" xorm:"create_time created"`
	DeleteTime time.Time `json:"deleteTime" xorm:"update_time deleted"`
}
