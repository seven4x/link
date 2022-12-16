package db

type UserVote struct {
	UserId int
	Id     int
	//t,l,c
	Type string
	//0,1,2
	IsLike int
}

type VoteInfo struct {
	Id       int
	Agree    int
	Disagree int
	Score    int
}
