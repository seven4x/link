package request

type NewTopicReq struct {
	Name       string `json:"name" validate:"required,min=2,max=140"`
	RefTopicId string `json:"refTopicId" validate:"required"`
	Position   int    `json:"position" validate:"required"`
	Predicate  string `json:"predicate"`
	Tags       string `json:"tags"`
}
