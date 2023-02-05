package job

import (
	"github.com/seven4x/link/app/log"
	"github.com/seven4x/link/db"
	"time"
)

func RefreshHotTopic() error {
	dao := db.NewDao()
	now := time.Now()
	d, _ := time.ParseDuration("-48h")
	expireD, _ := time.ParseDuration("52h")
	expireTime := now.Add(expireD)

	d2 := now.Add(d)
	res, err := dao.ListHotTopic(10, d2, now)
	if err != nil {
		log.Error(err.Error())
	}
	if len(res) == 0 {
		log.Info("no hot topic found.")
		return nil
	}

	topics := make([]*db.HotTopic, 0)
	for _, re := range res {
		topics = append(topics, &db.HotTopic{
			Id:         re,
			Expire:     expireTime,
			CreateTime: now,
		})
	}
	_, e := dao.Insert(&topics)

	return e
}
