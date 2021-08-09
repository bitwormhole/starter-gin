package main

import (
	"fmt"

	ginstarter "github.com/bitwormhole/starter-gin"
)

func main() {
	fmt.Println("src/main/go")
	ginstarter.InitGin().Run()
}
