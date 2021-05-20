package cron

import (
	logs "github.com/cihub/seelog"
	"github.com/robfig/cron/v3"

	"time"
)

func InitCron() {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	c := cron.New()
	c.AddFunc("@every 1s", func() {
		//logs.Info("启动第一次",time.Now().Format("2006-01-02 15:04:05"))
		//fmt.Println(tools.GetRunDir())
		logs.Error("启动")
	})
	c.Start()
	t1 := time.NewTimer(time.Second * 10)

	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
