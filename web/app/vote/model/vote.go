package model

type UserVote struct {
	UserId int
	Id     int
	//t,l,c
	Type rune
	//0,1,2
	IsLike rune
}

type VoteInfo struct {
	Id       int
	Agree    int
	DisAgree int
	Score    int
}
