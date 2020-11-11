package model

import (
	"time"
)

type Topic struct {
	Id        int64     `json:"id" xorm:"pk autoincr"`
	Name      string    `json:"name"`
	Tags      string    `json:"tags"`
	Lang      string    `json:"lang"`
	CreateBy  int       `json:"create_by"`
	Score     int       `json:"score"`
	Agree     int       `json:"agree"`
	Disagree  int       `json:"disagree"`
	CreatedAt time.Time `json:"create_time" xorm:"create_time created"`
	UpdateAt  time.Time `json:"update_time" xorm:"update_time updated"`
	DeletedAt time.Time `json:"-" xorm:"delete_time deleted"`
}

type TopicRel struct {
	Aid        int64     `json:"aid"`
	Bid        int64     `json:"bid"`
	Position   int       `json:"position"`
	CreateBy   string    `json:"create_by"`
	Predicate  string    `json:"predicate"`
	CreateTime time.Time `json:"create_time" xorm:"create_time created"`
	DeleteTime time.Time `json:"delete_time" xorm:"update_time deleted"`
}
