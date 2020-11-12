package request

type CreateLinkRequest struct {
	TopicId int    `validate:"required"`
	Title   string `validate:"required"`
	Url     string `validate:"required"`
	Comment string
	Group   string
	//1 web 2 chrome
	From int
}
