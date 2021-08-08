package main

import (
	startergin "github.com/bitwormhole/starter-gin/starter.gin"
)

func main() {
	app := startergin.InitGinApp()
	app.Use(startergin.ModuleDebug())
	app.Run()
}
