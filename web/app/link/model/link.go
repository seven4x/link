package model

import "time"

type Link struct {
	Id           int       `json:"id" xorm:"pk autoincr"`
	Link         string    `json:"link"`
	Title        string    `json:"title"`
	Group        string    `json:"group" xorm:"l_group"`
	Tags         string    `json:"tags" xorm:"varchar(140)"`
	From         int       `json:"from" xorm:"l_from char(1)"`
	TopicId      int       `json:"topicId"`
	Score        int       `json:"score"`
	Agree        int       `json:"agree"`
	Disagree     int       `json:"disagree"`
	CreateAt     time.Time `json:"create_time" xorm:"create_time created"`
	UpdateAt     time.Time `json:"update_time" xorm:"update_time updated"`
	DeletedAt    time.Time `json:"delete_time" xorm:"delete_time deleted"`
	CreateBy     int       `json:"create_by"`
	FirstComment string    `xorm:"-"`
}
