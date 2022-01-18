package factory

import (
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

// RestContext 配置REST上下文
type RestContext struct {
	markup.Component `id:"rest-web-context" class:"web-context"`

	Container      glass.Container    `inject:"#gin-web-container"`
	ControllerList []glass.Controller `inject:".rest-controller"`
	ContextPath    string             `inject:"${web.rest.context-path}"`
}

func (inst *RestContext) _Impl() glass.WebContext {
	return inst
}

// GetContainer 取WEB容器
func (inst *RestContext) GetContainer() glass.Container {
	return inst.Container
}

// GetContextPath 取上下文路径
func (inst *RestContext) GetContextPath() string {
	return inst.ContextPath
}

// GetControllers 取WEB控制器
func (inst *RestContext) GetControllers() []glass.Controller {
	return inst.ControllerList
}

////////////////////////////////////////////////////////////////////////////////

// StaticContext 配置静态文件上下文
type StaticContext struct {
	markup.Component `id:"static-web-context" class:"web-context"`

	Container      glass.Container    `inject:"#gin-web-container"`
	ControllerList []glass.Controller `inject:".static-web-controller"`
	ContextPath    string             `inject:"${web.static.context-path}"`
}

func (inst *StaticContext) _Impl() glass.WebContext {
	return inst
}

// GetContainer 取WEB容器
func (inst *StaticContext) GetContainer() glass.Container {
	return inst.Container
}

// GetContextPath 取上下文路径
func (inst *StaticContext) GetContextPath() string {
	return inst.ContextPath
}

// GetControllers 取WEB控制器
func (inst *StaticContext) GetControllers() []glass.Controller {
	return inst.ControllerList
}

////////////////////////////////////////////////////////////////////////////////
