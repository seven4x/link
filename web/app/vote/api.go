package vote

type VoteRequest struct {
	//topic,link,comment
	//t,l,c
	Type rune `validate:"oneof=t l c"`
	Id   int
	//0 不投票 1 喜欢 2 不喜欢
	//0,1,2
	IsLike   rune `validate:"oneof=0 1 2"`
	CreateBy int
}
