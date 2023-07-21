package main

import (
	"math/rand"
	"time"

	"github.com/xi123/libgo/core/base/sys/cmd"
	"github.com/xi123/libgo/core/base/task"
	"github.com/xi123/libgo/core/cb"
	"github.com/xi123/libgo/logs"
	"github.com/xi123/libgo/utils"
	"github.com/xi123/uploader/src/config"
	"github.com/xi123/uploader/src/global"
	"github.com/xi123/uploader/src/loader/handler"
	"github.com/xi123/uploader/src/loader/handler/sub"
	loader "github.com/xi123/uploader/src/loader/server"
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
	config.InitMonitorConfig(cmd.Conf())
	logs.SetTimezone(logs.Timezone(config.Config.Log.Monitor.Timezone))
	logs.SetMode(logs.Mode(config.Config.Log.Monitor.Mode))
	logs.SetStyle(logs.Style(config.Config.Log.Monitor.Style))
	logs.SetLevel(logs.Level(config.Config.Log.Monitor.Level))
	logs.Init(config.Config.Log.Monitor.Dir, global.Exe, 100000000)
	go func() {
		utils.ReadConsole(handler.OnInput)
	}()
	rand.Seed(time.Now().UnixNano())
	task.After(time.Duration(config.Config.Interval+(rand.Int()%50))*time.Second, cb.NewFunctor00(func() {
		handler.ReadConfig()
	}))
	loader.Run(cmd.Id(), config.ServiceName())
	sub.Start()
	sub.WaitAll()
	logs.Debugf("exit...")
	logs.Close()
}
