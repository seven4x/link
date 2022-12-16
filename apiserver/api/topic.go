package api

type CreateTopicRequest struct {
	Name       string `json:"name" validate:"required,min=2,max=140"`
	RefTopicId int    `json:"refId" validate:"required"`
	Position   int    `json:"position" validate:"oneof=1 2 3 4"`
	Predicate  string `json:"refDesc"`
	Tags       string `json:"tags"`
	Lang       string `json:"lang"`
	//1 公开编辑， 2个人 3 团队
	Scope int `json:"scope" validate:"oneof=1 2 3"`
}

type TopicDetail struct {
	Name       string `json:"name"`
	Id         int    `json:"id"`
	Icon       string `json:"icon"`
	CreateUser string `json:"createUser"`
	ShortCode  string `json:"shortCode"`
}

type SnapShot struct {
	Name string `json:"name"`
	//之所以使用字符串是因为Go int64的表示范围会超过JavaScript中number的表示范围
	Id        string `json:"id"`
	ShortCode string `json:"shortCode"`
}
