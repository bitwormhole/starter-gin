package glass

import (
	glassconf "github.com/bitwormhole/starter-gin/etc/glass/glass.conf"
	application "github.com/bitwormhole/starter/application"
)

const (
	myName     = "github.com/bitwormhole/starter-gin"
	myVersion  = "v0.0.8"
	myRevision = 8
)

// ExportModule 定义要导出的模块, use 【startergin.Module()】from outside
func ExportModule() application.Module {
	mod := &application.DefineModule{
		Name:     myName,
		Version:  myVersion,
		Revision: myRevision,
	}
	mod.OnMount = func(cb application.ConfigBuilder) error { return glassconf.MainConfig(cb, mod) }
	return mod
}
