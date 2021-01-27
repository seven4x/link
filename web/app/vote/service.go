package vote

import (
	"github.com/Seven4X/link/web/lib/api"
	"github.com/Seven4X/link/web/lib/api/messages"
)

type Service struct {
	dao *Dao
}

func NewService() *Service {
	return &Service{
		dao: NewDao(),
	}
}

/*
系统允不允许有爱又恨（即赞同又反对），不允许

  1 检查是否有历史投票记录
	投过票
		  投过赞成1->
			  1. 赞成 ，忽略
			  0，取消 ，更新成0, 累计-1, 赞成-1
			  2，反对 ，更新成2, 累计-2，赞成-1，反对+1
		  投过反对2->
			  1.赞成，更新1 ，累计+2，反对-1，赞成+1
			  0.取消  更新0 ，累计+1，反对-1
			  2。反对，忽略
		  投过0 ->
			  1.赞成，更新1 ，累计 +1，赞成+1
			  0.取消  忽略
			  2。反对， 更新2 ，累计-1，反对+1
  未投过票->
	  1  更新成1 ，累计+1,赞成+1
      2，取消投票（0），忽略
      3，投反对票，更新成2，累计投票 -1，反对+1

*/
func (s *Service) Vote(req *VoteRequest) (bool, *api.Err) {
	//查投票记录
	isLike, _ := s.dao.GetUserVote(req.CreateBy, req.Type, req.Id)
	if isLike == req.IsLike || (isLike == -1 && req.IsLike == '0') {
		return true, nil
	}

	session := s.dao.NewSession()
	session.Begin()

	//查询投票计分
	voteInfo, _ := GetVoteInfo(session, req.Type, req.Id)
	//重新计算计分
	dualScore(&voteInfo, isLike, req.IsLike)
	//更新计数
	err := UpdateVoteInfo(session, &voteInfo, req.Type)
	if err != nil {
		session.Rollback()
		return false, api.NewError(messages.GlobalErrorAboutDatabase)
	}
	uVote := UserVote{
		UserId: req.CreateBy,
		Id:     req.Id,
		Type:   req.Type,
		IsLike: req.IsLike,
	}
	//更新是否投票记录
	if isLike == -1 { //没有投票记录,新增
		err := CreateUserVote(session, &uVote)
		if err != nil {
			session.Rollback()
			return false, api.NewError(messages.GlobalErrorAboutDatabase)
		}
	} else {
		err := UpdateUserVote(session, &uVote)
		if err != nil {
			session.Rollback()
			return false, api.NewError(messages.GlobalErrorAboutDatabase)
		}
	}

	session.Commit()
	return false, nil
}

func (s *Service) ListIsLike(ids []interface{}, userId int, mtype string) (liked []UserVote, err error) {
	return s.dao.ListUserVoteByBusinessId(ids, userId, mtype)
}

func dualScore(info *VoteInfo, before rune, now rune) {
	if before != -1 {
		switch before {
		case '1':
			switch now {
			case '0':
				info.Score--
				info.Agree--
				break
			case '2':
				info.Score = info.Score - 2
				info.Agree--
				info.DisAgree++
				break
			}
			break
		case '2':
			switch now {
			case '0':
				info.Score++
				info.DisAgree--
				break
			case '1':
				info.Score = info.Score + 2
				info.Agree++
				info.DisAgree--
				break
			}
		case '0':
			switch now {
			case '1':
				info.Score++
				info.Agree++
			case '2':
				info.Score--
				info.DisAgree++
			}
		}
	} else {
		switch now {
		case '1':
			info.Score++
			info.Agree++
			break
		case '2':
			info.Score--
			info.DisAgree++
		}
	}
}
