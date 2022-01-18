package factory

import (
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
)

// DefaultContainer glass 的核心容器
type DefaultContainer struct {
	markup.Component `id:"gin-web-container"`

	Name         string                   `inject:"${server.name}"`
	IndexPages   string                   `inject:"${web.static.index-pages}"`
	Server       glass.Server             `inject:"#gin-web-server"`
	ContentTypes glass.ContentTypeManager `inject:"#gin-web-content-types"`
	Connectors   []glass.Connector        `inject:".web-server-connector"`
	Services     []glass.WebContext       `inject:".web-context"`
}

func (inst *DefaultContainer) _Impl() glass.Container {
	return inst
}

func (inst *DefaultContainer) GetName() string {
	return inst.Name
}

func (inst *DefaultContainer) GetServer() glass.Server {
	return inst.Server
}

func (inst *DefaultContainer) GetConnectors() []glass.Connector {
	return inst.Connectors
}

func (inst *DefaultContainer) GetContexts() []glass.WebContext {
	return inst.Services
}

func (inst *DefaultContainer) GetContentTypes() glass.ContentTypeManager {
	return inst.ContentTypes
}

func (inst *DefaultContainer) GetIndexPages() string {
	return inst.IndexPages
}
