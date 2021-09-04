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
	myVersion  = "v0.0.21"
	myRevision = 21
)

// Module 定义要导出的模块
func Module() application.Module {

	mod := &application.DefineModule{
		Name:     myName,
		Version:  myVersion,
		Revision: myRevision,
	}

	mod.OnMount = func(cb application.ConfigBuilder) error { return glassconf.MainConfig(cb, mod) }
	mod.Resources = srcmain.ExportResources()
	mod.AddDependency(starter.Module())

	return mod
}

// ModuleWithDevtools 定义要导出的模块(with devtools)
func ModuleWithDevtools() application.Module {

	mod := &application.DefineModule{
		Name:     myName + "/+devtools",
		Version:  myVersion,
		Revision: myRevision,
	}

	mod.OnMount = func(cb application.ConfigBuilder) error { return etcdevtools.ExportConfig(cb) }
	mod.Resources = devtools.ExportResources()
	mod.AddDependency(Module())

	return mod
}
