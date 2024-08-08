package common

import (
	"github.com/robfig/cron/v3"
	"sync"
)

var (
	_Cron *cron.Cron
)

func Cron() *cron.Cron {
	return _Cron
}

func InitCron() {
	sync.OnceFunc(func() {
		_Cron = cron.New(cron.WithSeconds())
		_Cron.Start()
	})()
}
