package config

import (
	"github.com/xi123/libgo/logs"
	"github.com/xi123/uploader/src/global"
)

func setServiceName(cb func(*IniConfig) string, c *IniConfig) {
	switch global.Name {
	case "":
		switch cb {
		case nil:
		default:
			global.Name = cb(c)
			logs.SetPrename(global.Name)
		}
	}
	switch global.Name == "" {
	case true:
		logs.Fatalf("error")
	}
}

func ServiceName() string {
	switch global.Name {
	case "":
		logs.Fatalf("error")
	}
	return global.Name
}
