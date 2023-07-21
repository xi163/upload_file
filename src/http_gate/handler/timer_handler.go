package handler

import (
	"time"

	"github.com/xi123/libgo/core/base/sys/cmd"
	"github.com/xi123/libgo/core/base/task"
	"github.com/xi123/libgo/core/cb"
	"github.com/xi123/uploader/src/config"
)

func ReadConfig() {
	config.ReadConfig(cmd.Conf())
	task.After(time.Duration(config.Config.Interval)*time.Second, cb.NewFunctor00(func() {
		ReadConfig()
	}))
}
