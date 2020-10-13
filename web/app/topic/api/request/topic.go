package request

type NewTopicReq struct {
	Name       string `json:"name" validate:"required,min=2,max=140"`
	RefTopicId int    `json:"refTopicId" validate:"required"`
	Position   int    `json:"position" validate:"oneof=1 2 3 4"`
	Predicate  string `json:"predicate"`
	Tags       string `json:"tags"`
}
