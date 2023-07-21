package http_gate

import (
	"net/http"

	"github.com/xi123/libgo/core/base/sys/cmd"
	"github.com/xi123/libgo/core/net/conn"
	"github.com/xi123/libgo/logs"
	"github.com/xi123/libgo/utils"
	"github.com/xi123/uploader/src/config"
	"github.com/xi123/uploader/src/global/httpsrv"
	"github.com/xi123/uploader/src/http_gate/handler"
)

// <summary>
// Router
// <summary>
type Router struct {
	server httpsrv.HttpServer
}

func (s *Router) Server() httpsrv.HttpServer {
	return s.server
}

func (s *Router) Init(id int, name string) {
}

func (s *Router) Run(id int, name string) {
	switch cmd.PatternArg("server") {
	case "":
		if id >= len(config.Config.HttpGate.Port) {
			logs.Fatalf("error id=%v HttpGate.Port.size=%v", id, len(config.Config.HttpGate.Port))
		}
		s.server = httpsrv.NewHttpServer(
			config.Config.HttpGate.Ip,
			config.Config.HttpGate.Port[id],
			config.Config.HttpGate.IdleTimeout)
	default:
		addr := conn.ParseAddress(cmd.PatternArg("server"))
		switch addr {
		case nil:
			logs.Fatalf("error")
		default:
			s.server = httpsrv.NewHttpServer(
				addr.Ip,
				utils.Atoi(addr.Port),
				config.Config.HttpGate.IdleTimeout)
		}
	}
	s.server.Router(config.Config.Path.UpdateCfg, s.UpdateConfigReq)
	s.server.Router(config.Config.Path.GetCfg, s.GetConfigReq)
	s.server.Router(config.Config.HttpGate.Path.Router, s.RouterReq)
	s.server.Run(id, name)
}

func (s *Router) UpdateConfigReq(w http.ResponseWriter, r *http.Request) {
	handler.UpdateCfgReq(w, r)
}

func (s *Router) GetConfigReq(w http.ResponseWriter, r *http.Request) {
	handler.GetCfgReq(w, r)
}

func (s *Router) RouterReq(w http.ResponseWriter, r *http.Request) {
	handler.RouterReq(w, r)
}
