package vote

import (
	"github.com/Seven4X/link/web/lib/log"
	"github.com/Seven4X/link/web/lib/store/db"
	"xorm.io/xorm"
)

type Dao struct {
	*xorm.Engine
}

func NewDao() (dao *Dao) {
	dao = &Dao{db.NewDb()}
	dao.NewSession()
	return
}

func (dao *Dao) GetUserVote(userId int, mtype string, mid int) (int, error) {
	var m UserVote
	b, err := dao.Where("user_id=?", userId).And("type=?", mtype).And("id=?", mid).Cols("is_like").Get(&m)
	if b {
		return m.IsLike, err
	} else {
		return -1, err
	}

}

func (dao *Dao) ListUserVoteByBusinessId(ids []interface{}, userId int, mtype string) ([]UserVote, error) {
	res := make([]UserVote, 0)
	err := dao.Table("user_vote").
		Cols("id", "is_like").
		In("id", ids...).
		And("user_id=?", userId).
		And("type=?", mtype).
		Find(&res)
	return res, err
}

func CreateUserVote(session *xorm.Session, uVote *UserVote) error {
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

func GetVoteInfo(session *xorm.Session, mtype string, mid int) (VoteInfo, error) {
	var result VoteInfo

	switch mtype {
	case "t":
		_, err := session.SQL("select score,agree,disagree,id from topic where id=?", mid).Get(&result)
		if err != nil {
			log.Error(err.Error())
		}
		return result, err
	case "l":
		_, err := session.SQL("select score,agree,disagree,id from link where id=?", mid).Get(&result)
		return result, err
	case "c":
		_, err := session.SQL("select score,agree,disagree,id from comment where id=?", mid).Get(&result)
		return result, err
	}
	return result, nil
}

func UpdateVoteInfo(session *xorm.Session, info *VoteInfo, mtype string) error {

	switch mtype {
	case "t":
		_, err := session.Exec("update topic set  score=?,agree=?,disagree=? where id=? ", info.Score, info.Agree, info.Disagree, info.Id)
		if err != nil {
			log.Error(err.Error())
			return err
		}
	case "l":
		_, err := session.Exec("update link set  score=?,agree=?,disagree=? where id=? ", info.Score, info.Agree, info.Disagree, info.Id)
		if err != nil {
			return err
		}

	case "c":
		_, err := session.Exec("update comment set  score=?,agree=?,disagree=? where id=? ", info.Score, info.Agree, info.Disagree, info.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateUserVote(session *xorm.Session, vote *UserVote) error {
	_, err := session.Exec("update user_vote set is_like=? where user_id=? and type=? and id=?", vote.IsLike, vote.UserId, vote.Type, vote.Id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
