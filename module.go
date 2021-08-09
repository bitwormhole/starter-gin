package ginstarter

import (
	"github.com/bitwormhole/starter"
	etcdevtools "github.com/bitwormhole/starter-gin/etc/devtools"
	glassconf "github.com/bitwormhole/starter-gin/etc/glass"
	"github.com/bitwormhole/starter-gin/src/devtools"
	srcmain "github.com/bitwormhole/starter-gin/src/main"
	"github.com/bitwormhole/starter/application"
)

const (
	myName     = "github.com/bitwormhole/starter-gin"
	myVersion  = "v0.0.8"
	myRevision = 8
)

// Module 定义要导出的模块
func Module() application.Module {

	dep1 := starter.Module()

	mod := &application.DefineModule{
		Name:     myName,
		Version:  myVersion,
		Revision: myRevision,
	}

	mod.OnMount = func(cb application.ConfigBuilder) error { return glassconf.MainConfig(cb, mod) }
	mod.Resources = srcmain.ExportResources()
	mod.Dependencies = []application.Module{dep1}

	return mod
}

// ModuleWithDevtools 定义要导出的模块(with devtools)
func ModuleWithDevtools() application.Module {

	res := devtools.ExportResources()

	dep1 := Module()
	deps := []application.Module{dep1}

	mod := &application.DefineModule{
		Name:         myName + "/+devtools",
		Version:      myVersion,
		Revision:     myRevision,
		Resources:    res,
		Dependencies: deps,
	}

	mod.OnMount = func(cb application.ConfigBuilder) error { return etcdevtools.ExportConfig(cb) }
	return mod
}
