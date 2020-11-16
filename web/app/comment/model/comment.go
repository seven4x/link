package model

type Comment struct {
	Id         int    `json:"id"`
	LinkId     int    `json:"link_id"`
	Context    string `json:"context"`
	Score      int    `json:"score"`
	Agree      int    `json:"agree"`
	Disagree   int    `json:"disagree"`
	CreateTime string `json:"create_time"`
}
