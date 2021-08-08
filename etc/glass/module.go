package glass

import (
	glassconf "github.com/bitwormhole/starter-gin/etc/glass/glass.conf"
	srcdebug "github.com/bitwormhole/starter-gin/src/debug"
	srcmain "github.com/bitwormhole/starter-gin/src/main"
	application "github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/starter"
)

const (
	myName     = "github.com/bitwormhole/starter-gin"
	myVersion  = "v0.0.8"
	myRevision = 8
)

// ExportModuleRelease 定义要导出的模块, use 【startergin.Module()】from outside
func ExportModuleRelease() application.Module {

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

// ExportModuleDebug 定义要导出的模块, use 【startergin.Module()】from outside
func ExportModuleDebug() application.Module {

	res := srcdebug.ExportResources()

	dep1 := ExportModuleRelease()
	deps := []application.Module{dep1}

	mod := &application.DefineModule{
		Name:         myName,
		Version:      myVersion + "-debug",
		Revision:     myRevision,
		Resources:    res,
		Dependencies: deps,
	}
	// mod.OnMount = func(cb application.ConfigBuilder) error { return glassconf.MainConfig(cb, mod) }
	mod.OnMount = func(cb application.ConfigBuilder) error { return nil }
	return mod
}
