package main

import (
	"embed"

	"github.com/bitwormhole/starter-gin/etc/demo"
	"github.com/bitwormhole/starter-gin/starter"
)

//go:embed src/main/resources
var resources embed.FS

func main() {
	appinit := starter.Gin().EmbedResources(&resources, "src/main/resources")
	appinit.Use(demo.ExportModule())
	appinit.Use(demo.ExportModule())
	appinit.Use(demo.ExportModule())
	appinit.Run()
}
