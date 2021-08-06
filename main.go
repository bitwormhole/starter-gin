package main

import (
	"embed"

	"github.com/bitwormhole/starter-gin/etc/demo"
	initd "github.com/bitwormhole/starter-gin/etc/init.d"
)

//go:embed src/main/resources
var resources embed.FS

func main() {

	appinit := initd.GinStarter()
	appinit.EmbedResources(&resources, "src/main/resources")

	// appinit.MountResources(initd.FileSystemResources("args:filesystem.resources"))

	appinit.Use(demo.ExportModule())
	appinit.Run()
}
