package main

import (
	"math/rand"
	"time"

	"github.com/xi123/libgo/core/base/sys/cmd"
	"github.com/xi123/libgo/core/base/task"
	"github.com/xi123/libgo/core/cb"
	"github.com/xi123/libgo/logs"
	"github.com/xi123/uploader/src/config"
	"github.com/xi123/uploader/src/global"
	"github.com/xi123/uploader/src/http_gate/handler"
	http_gate "github.com/xi123/uploader/src/http_gate/server"
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
	config.InitHttpGateConfig(cmd.Conf())
	logs.SetTimezone(logs.Timezone(config.Config.Log.HttpGate.Timezone))
	logs.SetMode(logs.Mode(config.Config.Log.HttpGate.Mode))
	logs.SetStyle(logs.Style(config.Config.Log.HttpGate.Style))
	logs.SetLevel(logs.Level(config.Config.Log.HttpGate.Level))
	logs.Init(config.Config.Log.HttpGate.Dir, global.Exe, 100000000)
	rand.Seed(time.Now().UnixNano())
	task.After(time.Duration(config.Config.Interval+(rand.Int()%50))*time.Second, cb.NewFunctor00(func() {
		handler.ReadConfig()
	}))
	http_gate.Run(cmd.Id(), config.ServiceName())
	logs.Close()
}
