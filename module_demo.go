package ginstarter

import (
	"embed"

	"github.com/bitwormhole/starter-gin/gen/config_demo"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

//go:embed src/demo/resources
var theDemoResFS embed.FS

// ModuleForDemo 定义要导出的模块(with demo)
func ModuleForDemo() application.Module {

	parent := ModuleWithDevtools()
	mb := application.ModuleBuilder{}

	mb.Name(parent.GetName() + "/+demo")
	mb.Version(parent.GetVersion())
	mb.Revision(parent.GetRevision())

	mb.Resources(collection.LoadEmbedResources(&theDemoResFS, "src/demo/resources"))
	mb.OnMount(config_demo.ExportConfigForStarterGinDemo)
	mb.Dependency(parent)

	return mb.Create()
}
