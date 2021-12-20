package job

import (
	"github.com/seven4x/link/web/topic"
	"github.com/seven4x/link/web/util"
	"time"
)

func RefreshHotTopic() error {
	dao := topic.NewDao()
	now := time.Now()
	d, _ := time.ParseDuration("-48h")
	expireD, _ := time.ParseDuration("52h")
	expireTime := now.Add(expireD)

	d2 := now.Add(d)
	res, err := dao.ListHotTopic(10, d2, now)
	if err != nil {
		util.Error(err.Error())
	}
	if len(res) == 0 {
		util.Info("no hot topic found.")
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