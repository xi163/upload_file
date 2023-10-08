package main

import (
	"math/rand"
	"time"

	"github.com/cwloo/gonet/core/base/sys/cmd"
	"github.com/cwloo/gonet/core/base/task"
	"github.com/cwloo/gonet/core/cb"
	"github.com/cwloo/gonet/logs"
	"github.com/cwloo/uploader/src/config"
	"github.com/cwloo/uploader/src/file_client/handler"
	"github.com/cwloo/uploader/src/global"
)

func init() {
	cmd.InitArgs(func(arg *cmd.ARG) {
		arg.SetConf("config/conf.ini")
	})
}

func main() {
	cmd.ParseArgs()
	config.InitClientConfig(cmd.Conf())
	logs.SetTimezone(logs.Timezone(config.Config.Log.Client.Timezone))
	logs.SetMode(logs.Mode(config.Config.Log.Client.Mode))
	logs.SetStyle(logs.Style(config.Config.Log.Client.Style))
	logs.SetLevel(logs.Level(config.Config.Log.Client.Level))
	logs.Init(config.Config.Log.Client.Dir, global.Exe, 100000000)
	rand.Seed(time.Now().UnixNano())
	task.After(time.Duration(config.Config.Interval+(rand.Int()%50))*time.Second, cb.NewFunctor00(func() {
		handler.ReadConfig()
	}))
	switch config.Config.Client.Upload.MultiFile > 0 {
	case true:
	default:
		handler.Upload()
	}
}
