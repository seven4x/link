package request

type ListLinkRequest struct {
	Prev   int
	Size   int
	Tid    int `validate:"required"`
	Group  string
	UserId int
}
