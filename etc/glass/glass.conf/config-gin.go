package glassconf

import (
	"github.com/bitwormhole/starter-gin/glass"
	srcdebuggolang "github.com/bitwormhole/starter-gin/src/debug/golang"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

type theContainer struct {
	markup.Component
	instance *glass.Container `id:"gin-web-container"`

	AppContext   application.Context      `inject:"context"`
	Name         string                   `inject:"${server.name}"`
	Server       *glass.Server            `inject:"#gin-web-server"`
	Connectors   []glass.Connector        `inject:".web-server-connector"`
	Services     []*glass.WebContext      `inject:".web-context"`
	ContentTypes glass.ContentTypeManager `inject:"#gin-web-content-types"`
	IndexPages   string                   `inject:"${web.static.index-pages}"`
}

type theServer struct {
	markup.Component
	instance *glass.Server `id:"gin-web-server" class:"looper" initMethod:"Start" destroyMethod:"Stop"`

	Container *glass.Container `inject:"#gin-web-container"`
}

type theHTTPConnector struct {
	markup.Component
	instance *glass.HTTPConnector `class:"web-server-connector"`

	GinMode string `inject:"${gin.mode}"`

	Host   string `inject:"${server.host}"`
	Port   int    `inject:"${server.port}"`
	Enable bool   `inject:"${server.enable}"`
}

type theHTTPSConnector struct {
	markup.Component
	instance *glass.HTTPSConnector `class:"web-server-connector"`

	GinMode string `inject:"${gin.mode}"`

	Host   string `inject:"${server.https.host}"`
	Port   int    `inject:"${server.https.port}"`
	Enable bool   `inject:"${server.https.enable}"`

	CertificateFile string `inject:"${server.https.cert-file}"`
	PrivateKeyFile  string `inject:"${server.https.key-file}"`
}

type theStaticWebService struct {
	markup.Component
	instance *glass.WebContext `class:"web-context"`

	Container   *glass.Container   `inject:"#gin-web-container"`
	ContextPath string             `inject:"${web.static.context-path}"`
	Controllers []glass.Controller `inject:".static-web-controller"`
}

type theRestWebService struct {
	markup.Component
	instance *glass.WebContext `class:"web-context"`

	Container   *glass.Container   `inject:"#gin-web-container"`
	ContextPath string             `inject:"${web.rest.context-path}"`
	Controllers []glass.Controller `inject:".rest-controller"`
}

type theStaticWebController struct {
	markup.Component
	instance *glass.StaticWebResourcesController `class:"static-web-controller"`

	Root      string           `inject:"${web.static.root}"`
	Container *glass.Container `inject:"#gin-web-container"`
}

type theWebDevtoolsController struct {
	markup.Component
	instance *srcdebuggolang.DevtoolsController `class:"rest-controller"`

	AppContext application.Context `inject:"context"`
	Enable     bool                `inject:"${web.devtools.enable}"`
}

type theWebContentTypes struct {
	markup.Component
	instance *glass.DefaultContentTypeManager `id:"gin-web-content-types"`

	AppContext      application.Context `inject:"context"`
	TypesProperties string              `inject:"${web.static.content-types-properties}"`
}

type theWebErrorController struct {
	markup.Component
	instance *glass.ErrorController `class:"static-web-controller"`

	ResourcePath string `inject:"${web.error-page.resource}"`
	ContentType  string `inject:"${web.error-page.content-type}"`
	Status       int    `inject:"${web.error-page.status}"`

	Container *glass.Container `inject:"#gin-web-container"`
}
