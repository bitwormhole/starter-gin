package glass

import application "github.com/bitwormhole/starter/application"

// Module 对外导出本模块
func Module() application.Module {
	return &application.DefineModule{
		Name:     "github.com/bitwormhole/starter-gin/glass",
		Version:  "v0.0.5",
		Revision: 5,
		OnMount:  func(cb application.ConfigBuilder) error { return mainConfig(cb) },
	}
}
