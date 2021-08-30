package srcmain

import (
	"embed"

	"github.com/bitwormhole/starter/collection"
)

//go:embed resources
var resourcesDir embed.FS

// ExportResources 导出资源组
func ExportResources() collection.Resources {
	return collection.LoadEmbedResources(&resourcesDir, "resources")
}
