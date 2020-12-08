package request

type NewCommentRequest struct {
	Content string `validate:"max=140"`
	LinkId  int
	//
	CreateBy int
}
