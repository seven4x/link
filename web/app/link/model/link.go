package model

type Link struct {
	ID         int    `json:"id"`
	Link       string `json:"link"`
	Title      string `json:"title"`
	Topic      int    `json:"topic"`
	ShowCommit string `json:"show_commit"`
	Score      int    `json:"score"`
	Agree      int    `json:"agree"`
	Disagree   int    `json:"disagree"`
	CreateTime string `json:"create_time"`
}
