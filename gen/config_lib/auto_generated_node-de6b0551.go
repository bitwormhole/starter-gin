// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package config_lib

import (
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	factory0x58e669 "github.com/bitwormhole/starter-gin/glass/factory"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComDefaultContentTypeManager struct {
	instance *glass0x47343f.DefaultContentTypeManager
	 markup0x23084a.Component `id:"gin-web-content-types"`
	AppContext application0x67f6c5.Context `inject:"context"`
	TypesProperties string `inject:"${web.static.content-types-properties}"`
}


type pComErrorController struct {
	instance *glass0x47343f.ErrorController
	 markup0x23084a.Component `class:"static-web-controller"`
	ResourcePath string `inject:"${web.error-page.resource}"`
	ContentType string `inject:"${web.error-page.content-type}"`
	Status int `inject:"${web.error-page.status}"`
	Container glass0x47343f.Container `inject:"#gin-web-container"`
	Context application0x67f6c5.Context `inject:"context"`
}


type pComContextBindController struct {
	instance *factory0x58e669.ContextBindController
	 markup0x23084a.Component `class:"rest-controller"`
	Order int `inject:"${webfilter.context.order}"`
}


type pComHTTPSConnector struct {
	instance *factory0x58e669.HTTPSConnector
	 markup0x23084a.Component `class:"web-server-connector"`
	GinMode string `inject:"${gin.mode}"`
	Host string `inject:"${server.https.host}"`
	Port int `inject:"${server.https.port}"`
	MyEnable bool `inject:"${server.https.enable}"`
	CertificateFile string `inject:"${server.https.cert-file}"`
	PrivateKeyFile string `inject:"${server.https.key-file}"`
}


type pComHTTPConnector struct {
	instance *factory0x58e669.HTTPConnector
	 markup0x23084a.Component `class:"web-server-connector"`
	GinMode string `inject:"${gin.mode}"`
	Host string `inject:"${server.host}"`
	Port int `inject:"${server.port}"`
	MyEnable bool `inject:"${server.enable}"`
}


type pComWebContainer struct {
	instance *factory0x58e669.WebContainer
	 markup0x23084a.Component `id:"gin-web-container" class:"life"`
	Name string `inject:"${server.name}"`
	IndexPages string `inject:"${web.static.index-pages}"`
	ContentTypes glass0x47343f.ContentTypeManager `inject:"#gin-web-content-types"`
	Connectors []glass0x47343f.Connector `inject:".web-server-connector"`
	Contexts []glass0x47343f.WebContext `inject:".web-context"`
}


type pComRestContext struct {
	instance *factory0x58e669.RestContext
	 markup0x23084a.Component `id:"rest-web-context" class:"web-context"`
	Container glass0x47343f.Container `inject:"#gin-web-container"`
	ControllerList []glass0x47343f.Controller `inject:".rest-controller"`
	ContextPath string `inject:"${web.rest.context-path}"`
}


type pComStaticContext struct {
	instance *factory0x58e669.StaticContext
	 markup0x23084a.Component `id:"static-web-context" class:"web-context"`
	Container glass0x47343f.Container `inject:"#gin-web-container"`
	ControllerList []glass0x47343f.Controller `inject:".static-web-controller"`
	ContextPath string `inject:"${web.static.context-path}"`
}


type pComStaticWebResourcesController struct {
	instance *glass0x47343f.StaticWebResourcesController
	 markup0x23084a.Component `class:"static-web-controller"`
	Root string `inject:"${web.static.root}"`
	Container glass0x47343f.Container `inject:"#gin-web-container"`
	Context application0x67f6c5.Context `inject:"context"`
}

