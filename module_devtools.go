package ginstarter

import (
	"embed"

	"github.com/bitwormhole/starter-gin/gen/config_devtools"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

//go:embed src/devtools/resources
var theDevtoolsResFS embed.FS

// ModuleWithDevtools 定义要导出的模块(with devtools)
func ModuleWithDevtools() application.Module {

	mb := &application.ModuleBuilder{}
	mb.Name(myName + "/+devtools").Version(myVersion).Revision(myRevision)
	mb.Dependency(Module())
	mb.Resources(collection.LoadEmbedResources(&theDevtoolsResFS, "src/devtools/resources"))

	// mod.OnMount = func(cb application.ConfigBuilder) error { return etcdevtools.ExportConfig(cb) }
	mb.OnMount(config_devtools.ExportConfigForGinStarterDevtools)

	return mb.Create()
}
