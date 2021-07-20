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

	// eng := gin.Default()
	// eng.GET("", func(c *gin.Context) {})
	// eng.Use(func(c *gin.Context) {})

	res := config.CreateEmbedFsResources(&resources, "src/main/resources")
	err := demo.RunDemo(os.Args, res)
	if err != nil {
		panic(err)
	}
}
