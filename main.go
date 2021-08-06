package main

import (
	"embed"

	initd "github.com/bitwormhole/starter-gin/etc/init.d"
)

//go:embed src/main/resources
var resources embed.FS

func main() {

	appinit := initd.GinStarter()

	appinit.MountResources(initd.FileSystemResources("args:filesystem.resources"))

	// appinit.EmbedResources(&resources, "src/main/resources")

	// appinit.Use(demo.ExportModule())
	// appinit.Use(demo.ExportModule())
	// appinit.Use(demo.ExportModule())

	appinit.Run()
}
