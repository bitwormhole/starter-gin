package main

import (
	"embed"

	"github.com/bitwormhole/starter-gin/starter"
)

//go:embed src/main/resources
var resources embed.FS

func main() {
	starter.Gin().EmbedResources(&resources, "src/main/resources").Run()
}
