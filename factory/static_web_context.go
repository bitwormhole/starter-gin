package factory

import (
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
)

// StaticWebContext 是静态web内容的上下文
type StaticWebContext struct {
	markup.Component `class:"web-context"`

	ContextPath string             `inject:"${web.static.context-path}"`
	Container   glass.Container    `inject:"#gin-web-container"`
	Controllers []glass.Controller `inject:".static-web-controller"`
}

func (inst *StaticWebContext) _Impl() glass.WebContext {
	return inst
}

func (inst *StaticWebContext) GetContainer() glass.Container {
	return inst.Container
}

func (inst *StaticWebContext) GetContextPath() string {
	return inst.ContextPath
}

func (inst *StaticWebContext) GetControllers() []glass.Controller {
	return inst.Controllers
}
