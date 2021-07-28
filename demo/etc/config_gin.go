package etcdemo

import (
	"github.com/bitwormhole/starter-gin/demo/elements"
	"github.com/bitwormhole/starter-gin/web"
	application "github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

type theGinHandlerExample1 struct {
	markup.Component
	instance *elements.ExampleGinController
}

type the404handler struct {
	markup.Component
	instance *web.Http404pageController
}

type theStaticHandler struct {
	markup.Component
	instance *web.StaticResourcesController

	ApplicationContext application.Context `inject:"context"`
}
