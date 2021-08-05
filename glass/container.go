package glass

import "github.com/bitwormhole/starter/application"

// Container 表示一个web容器(gin-glass)
type Container struct {

	// public
	Name         string
	Server       *Server
	Connectors   []Connector
	Services     []*WebContext
	ContentTypes ContentTypeManager
	IndexPages   string
	AppContext   application.Context

	// private
	runtime *serverRuntime
}
