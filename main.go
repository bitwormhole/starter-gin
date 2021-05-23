package main

import (
	"embed"
	"os"

	"github.com/bitwormhole/starter-gin/demo"
	"github.com/bitwormhole/starter/application/config"
)

//go:embed src/main/resources
var resources embed.FS

func main() {
	res := config.CreateEmbedFsResources(&resources, "src/main/resources")
	err := demo.RunDemo(os.Args, res)
	if err != nil {
		panic(err)
	}
}
