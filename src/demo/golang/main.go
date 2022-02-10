package main

import (
	"github.com/bitwormhole/starter"
	ginstarter "github.com/bitwormhole/starter-gin"
)

func main() {
	i := starter.InitApp()
	i.Use(ginstarter.ModuleForDemo())
	i.Run()
}
