package etcdemo

import (
	"github.com/bitwormhole/starter-gin/demo/elements"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

func Config(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}

type theGinHandlerExample1 struct {
	markup.Component
	instance *elements.ExampleGinController
}
