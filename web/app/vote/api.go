package vote

type VoteRequest struct {
	//topic,link,comment
	//t,l,c
	TypeCode string `validate:"oneof=topic link comment"`
	Type     string
	Id       int
	//0 不投票 1 喜欢 2 不喜欢
	//0,1,2
	IsLike   int `validate:"oneof=0 1 2"`
	CreateBy int
}
