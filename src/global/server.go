package global

import "github.com/xi123/libgo/core/net/tcp/tcpserver"

// <summary>
// TCPServer
// <summary>
type TCPServer interface {
	Server() tcpserver.TCPServer
	Init(id int, name string)
	Run(id int, name string)
}
