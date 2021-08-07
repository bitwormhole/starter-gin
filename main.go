package main

import (
	"embed"

	"github.com/bitwormhole/starter-gin/etc/demo"
	startergin "github.com/bitwormhole/starter-gin/starter.gin"
)

//go:embed src/main/resources
var resources embed.FS

func main() {
	app := startergin.InitGinApp()
	app.EmbedResources(&resources, "src/main/resources")
	app.Use(demo.ExportModule())
	app.Run()
}
