package glass

import (
	glassconf "github.com/bitwormhole/starter-gin/etc/glass/glass.conf"
	application "github.com/bitwormhole/starter/application"
)

// ExportModule 定义要导出的模块, use 【startergin.Module()】from outside
func ExportModule() application.Module {
	return &application.DefineModule{
		Name:     "github.com/bitwormhole/starter-gin",
		Version:  "v0.0.7",
		Revision: 7,
		OnMount:  func(cb application.ConfigBuilder) error { return glassconf.MainConfig(cb) },
	}
}
