package model

import "time"

type Topic struct {
	Id        int       `json:"id" xorm:"pk autoincr"`
	Name      string    `json:"name"`
	Tags      string    `json:"tags"`
	CreateBy  string    `json:"create_by"`
	Score     int       `json:"score"`
	Agree     int       `json:"agree"`
	Disagree  int       `json:"disagree"`
	CreatedAt time.Time `json:"create_time" xorm:"created"`
	UpdateAt  time.Time `json:"update_time" xorm:"updated"`
	DeletedAt time.Time `json:"-" xorm:"deleted"`
}

type TopicRel struct {
	Aid        int       `json:"aid"`
	Bid        int       `json:"bid"`
	Position   int       `json:"position"`
	CreateBy   string    `json:"create_by"`
	Predicate  string    `json:"predicate"`
	CreateTime time.Time `json:"create_time" xorm:"created"`
	DeleteTime time.Time `json:"delete_time" xorm:"deleted"`
}
