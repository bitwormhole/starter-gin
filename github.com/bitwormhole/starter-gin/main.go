package main

import (
	"fmt"

	"github.com/bitwormhole/starter-gin/etc/startergin"
)

func main() {
	cfg := startergin.Config()
	fmt.Println("hello, this is ", cfg)
}
