package factory

import (
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
)

// RESTfulWebContext 是RESTful web内容的上下文
type RESTfulWebContext struct {
	markup.Component `class:"web-context"`

	ContextPath string             `inject:"${web.rest.context-path}"`
	Container   glass.Container    `inject:"#gin-web-container"`
	Controllers []glass.Controller `inject:".rest-controller"`
}

func (inst *RESTfulWebContext) _Impl() glass.WebContext {
	return inst
}

func (inst *RESTfulWebContext) GetContainer() glass.Container {
	return inst.Container
}

func (inst *RESTfulWebContext) GetContextPath() string {
	return inst.ContextPath
}

func (inst *RESTfulWebContext) GetControllers() []glass.Controller {
	return inst.Controllers
}
