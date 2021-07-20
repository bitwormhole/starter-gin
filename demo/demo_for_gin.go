package demo

import (
	"fmt"

	etc_demo "github.com/bitwormhole/starter-gin/demo/etc"
	etc_gin "github.com/bitwormhole/starter-gin/etc/gin"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/collection"
)

func RunDemo(args []string, res collection.Resources) error {

	configurer := config.NewBuilder()
	configurer.SetResources(res)

	// etc_gin.Config(configurer)
	// etc_demo.Config(configurer)

	etc_gin.Config(configurer)
	etc_demo.Config(configurer)

	configuration := configurer.Create()
	context, err := application.Run(configuration, args)
	if err != nil {
		return err
	}

	err = application.Loop(context)
	if err != nil {
		return err
	}

	code, err := application.Exit(context)
	if err != nil {
		return err
	}

	fmt.Println("Exit with code ", code)
	return nil
}
