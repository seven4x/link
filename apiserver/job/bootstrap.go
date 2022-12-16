package job

import (
	"github.com/robfig/cron/v3"
	"github.com/seven4x/link/app"
	"time"
)

func StartJob() *cron.Cron {
	local, _ := time.LoadLocation("Local")
	c := cron.New(cron.WithLocation(local))
	c.AddFunc("@midnight", func() {
		err := RefreshHotTopic()
		if err != nil {
			app.Error(err.Error())
		}
	})
	c.AddFunc("@hourly", func() {
		app.DumpCuckooFilter()
	})
	//c.AddFunc("@every 10s", func() {
	//	log.Info("活")
	//})
	c.Start()

	return c
}
