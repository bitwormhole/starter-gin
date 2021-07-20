package etcdemo

import (
	"github.com/bitwormhole/starter-gin/demo/elements"
	"github.com/bitwormhole/starter/markup"
)

type theGinHandlerExample1 struct {
	markup.Component
	instance *elements.ExampleGinController
}
