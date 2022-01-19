package ginstarter

import (
	"embed"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter-gin/gen/config_lib"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myName     = "github.com/bitwormhole/starter-gin"
	myVersion  = "v0.1.0"
	myRevision = 25
)

//go:embed src/main/resources
var theMainResFS embed.FS

// Module 定义要导出的模块
func Module() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(myName).Version(myVersion).Revision(myRevision)
	mb.Resources(collection.LoadEmbedResources(&theMainResFS, "src/main/resources"))
	mb.Dependency(starter.Module())

	// mb.OnMount = func(cb application.ConfigBuilder) error { return glassconf.MainConfig(cb, mod) }
	mb.OnMount(config_lib.ExportConfigForGinStarterLib)

	return mb.Create()
}
