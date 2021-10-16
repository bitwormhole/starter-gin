package ginstarter

import (
	"embed"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter-gin/src/main/etc"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

const (
	myName     = "github.com/bitwormhole/starter-gin"
	myVersion  = "v0.0.24"
	myRevision = 24
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
	mb.OnMount(etc.ExportGinConfig)

	return mb.Create()
}
