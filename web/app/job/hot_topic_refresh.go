package job

import (
	"github.com/Seven4X/link/web/app/topic"
	"github.com/Seven4X/link/web/lib/log"
	"time"
)

func RefreshHotTopic() error {
	dao := topic.NewDao()
	now := time.Now()
	d, _ := time.ParseDuration("-48h")
	expireD, _ := time.ParseDuration("24h")
	expireTime := now.Add(expireD)

	d2 := now.Add(d)
	res, err := dao.ListHotTopic(3, d2, now)
	if err != nil {
		log.Error(err.Error())
	}
	if len(res) == 0 {
		log.Info("no hot topic found.")
		return nil
	}

	topics := make([]*topic.HotTopic, 0)
	for _, re := range res {
		topics = append(topics, &topic.HotTopic{
			Id:         re,
			Expire:     expireTime,
			CreateTime: now,
		})
	}
	_, e := dao.Insert(&topics)

	return e
}
