package dao

import (
	"github.com/Seven4X/link/web/app/vote/model"
	"github.com/Seven4X/link/web/library/log"
	"github.com/Seven4X/link/web/library/store/db"
	"github.com/xormplus/xorm"
)

type Dao struct {
	*xorm.Engine
}

func New() (dao *Dao) {
	dao = &Dao{db.NewDb()}
	dao.NewSession()
	return
}

func (dao *Dao) GetUserVote(userId int, mtype rune, mid int) rune {
	var m model.UserVote
	b, _ := dao.Where("user_id=?", userId).And("type=?", mtype).And("id=?", mid).Cols("is_like").Get(&m)
	if b {
		return m.IsLike
	} else {
		return -1
	}

}

func CreateUserVote(session *xorm.Session, uVote *model.UserVote) error {
	id, err := session.InsertOne(uVote)
	if err != nil {
		log.Errorw("CreateUserVote",
			"uid", uVote.UserId,
			"id", uVote.Id,
			"type", uVote.Type,
			"id", id, "err", err.Error())
	}

	return err

}
func GetVoteInfo(session *xorm.Session, mtype rune, mid int) model.VoteInfo {
	var result model.VoteInfo

	switch mtype {
	case 't':
		_, err := session.SQL("select score,agree,disagree,id from topic where id=?", mid).Get(&result)
		if err != nil {
			log.Error(err.Error())
		}
		return result
	case 'l':
		session.SQL("select score,agree,disagree,id from link where id=?", mid).Get(&result)
		return result
	case 'c':
		session.SQL("select score,agree,disagree,id from comment where id=?", mid).Get(&result)
		return result
	}
	return result
}

func UpdateVoteInfo(session *xorm.Session, info *model.VoteInfo, mtype rune) error {

	switch mtype {
	case 't':
		_, err := session.Exec("update topic set  score=?,agree=?,disagree=? where id=? ", info.Score, info.Agree, info.DisAgree, info.Id)
		if err != nil {
			log.Error(err.Error())
			return err
		}
	case 'l':
		_, err := session.Exec("update link set  score=?,agree=?,disagree=? where id=? ", info.Score, info.Agree, info.DisAgree, info.Id)
		if err != nil {
			return err
		}

	case 'c':
		_, err := session.Exec("update comment set  score=?,agree=?,disagree=? where id=? ", info.Score, info.Agree, info.DisAgree, info.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateUserVote(session *xorm.Session, vote *model.UserVote) error {
	_, err := session.Exec("update user_vote set is_like=? where user_id=? and type=? and id=?", vote.IsLike, vote.UserId, vote.Type, vote.Id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
