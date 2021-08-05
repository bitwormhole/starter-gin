package demo

import (
	"github.com/bitwormhole/starter/application"
)

func ExportModule() application.Module {
	return &application.DefineModule{

		Name:     "glass/demo",
		Version:  "0.1",
		Revision: 1,

		OnMount: func(cb application.ConfigBuilder) error { return mainConfig(cb) },
	}
}

func mainConfig(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
