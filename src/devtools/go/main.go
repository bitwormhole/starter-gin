package main

import (
	"fmt"

	ginstarter "github.com/bitwormhole/starter-gin"
)

func main() {
	fmt.Println("src/devtools/go")
	i := ginstarter.InitGin()
	i.Use(ginstarter.ModuleWithDevtools())
	i.Run()
}
