package glass

// Container 表示一个web容器(gin-glass)
type Container interface {

	// GetServer() Server                   //   `inject:"#gin-web-server"`

	GetName() string                     //  `inject:"${server.name}"`
	GetConnectors() []Connector          //  `inject:".web-server-connector"`
	GetContexts() []WebContext           //  `inject:".web-context"`
	GetContentTypes() ContentTypeManager //  `inject:"#gin-web-content-types"`
	GetIndexPages() string               //  `inject:"${web.static.index-pages}"`
	GetIndexPageNames() []string         //  `inject:"${web.static.index-pages}"`

	// markup.Component `id:"gin-web-container"`

	// // public

	// AppContext   application.Context `inject:"context"`
	// Name         string              `inject:"${server.name}"`
	// Server       *Server             `inject:"#gin-web-server"`
	// Connectors   []Connector         `inject:".web-server-connector"`
	// Services     []*WebContext       `inject:".web-context"`
	// ContentTypes ContentTypeManager  `inject:"#gin-web-content-types"`
	// IndexPages   string              `inject:"${web.static.index-pages}"`

	// // private

	// runtime *serverRuntime

	Shutdown() error
}
