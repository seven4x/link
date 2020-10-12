package topic

type Topic struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Tags       string `json:"tags"`
	CreateBy   string `json:"create_by"`
	Score      int    `json:"score"`
	Agree      int    `json:"agree"`
	Disagree   int    `json:"disagree"`
	CreateTime string `json:"create_time"`
}
