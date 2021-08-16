package srctest

import (
	"embed"

	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/collection"
)

//go:embed resources
var resourcesDir embed.FS

// ExportResources 导出资源
func ExportResources() collection.Resources {
	return config.LoadResourcesFromEmbedFS(&resourcesDir, "resources")
}
