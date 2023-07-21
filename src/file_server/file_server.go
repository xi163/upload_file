package main

import (
	"math/rand"
	"time"

	"github.com/xi123/libgo/core/base/sys/cmd"
	"github.com/xi123/libgo/core/base/task"
	"github.com/xi123/libgo/core/cb"
	"github.com/xi123/libgo/logs"
	"github.com/xi123/uploader/src/config"
	"github.com/xi123/uploader/src/file_server/handler"
	file_server "github.com/xi123/uploader/src/file_server/server"
	"github.com/xi123/uploader/src/global"
)

func init() {
	cmd.InitArgs(func(arg *cmd.ARG) {
		arg.SetConf("config/conf.ini")
		arg.AppendPattern("server", "server", "srv", "svr", "s")
		arg.AppendPattern("rpc", "rpc", "r")
	})
}

func main() {
	cmd.ParseArgs()
	config.InitFileConfig(cmd.Conf())
	logs.SetTimezone(logs.Timezone(config.Config.Log.File.Timezone))
	logs.SetMode(logs.Mode(config.Config.Log.File.Mode))
	logs.SetStyle(logs.Style(config.Config.Log.File.Style))
	logs.SetLevel(logs.Level(config.Config.Log.File.Level))
	logs.Init(config.Config.Log.File.Dir, global.Exe, 100000000)

	task.After(time.Duration(config.Config.File.Upload.PendingTimeout)*time.Second, cb.NewFunctor00(func() {
		handler.PendingUploader()
	}))

	task.After(time.Duration(config.Config.File.Upload.FileExpiredTimeout)*time.Second, cb.NewFunctor00(func() {
		handler.ExpiredFile()
	}))
	rand.Seed(time.Now().UnixNano())
	task.After(time.Duration(config.Config.Interval+(rand.Int()%50))*time.Second, cb.NewFunctor00(func() {
		handler.ReadConfig()
	}))
	file_server.Run(cmd.Id(), config.ServiceName())
	logs.Close()
}
