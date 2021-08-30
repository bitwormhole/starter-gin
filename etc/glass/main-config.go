package glass

import (
	"strconv"

	"github.com/bitwormhole/starter/application"
)

const (
	// WebStaticContextPath  静态web资源的上下文路径
	WebStaticContextPath            = "web.static.context-path"
	WebStaticIndexPages             = "web.static.index-pages"
	WebStaticContentTypesProperties = "web.static.content-types-properties"
	WebStaticRoot                   = "web.static.root"

	// WebRestContextPath  REST-web资源的上下文路径
	WebRestContextPath = "web.rest.context-path"

	ServerName   = "server.name"
	ServerPort   = "server.port"
	ServerHost   = "server.host"
	ServerEnable = "server.enable"

	ServerHttpsName     = "server.https.name"
	ServerHttpsPort     = "server.https.port"
	ServerHttpsHost     = "server.https.host"
	ServerHttpsEnable   = "server.https.enable"
	ServerHttpsKeyFile  = "server.https.key-file"
	ServerHttpsCertFile = "server.https.cert-file"

	WebErrorPageResource    = "web.error-page.resource"
	WebErrorPageContentType = "web.error-page.content-type"
	WebErrorPageStatus      = "web.error-page.status"

	WebDevToolsEnable = "web.devtools.enable"

	GinMode = "gin.mode"
)

func MainConfig(cb application.ConfigBuilder, module application.Module) error {

	p := cb.DefaultProperties()

	p.SetProperty("module.starter-gin.name", module.GetName())
	p.SetProperty("module.starter-gin.version", module.GetVersion())
	p.SetProperty("module.starter-gin.revision", strconv.Itoa(module.GetRevision()))

	p.SetProperty(WebRestContextPath, "/api")
	p.SetProperty(WebDevToolsEnable, "true")

	p.SetProperty(WebStaticContextPath, "/")
	p.SetProperty(WebStaticIndexPages, "index.html, index.htm")
	p.SetProperty(WebStaticContentTypesProperties, "res:///content-types.properties")
	p.SetProperty(WebStaticRoot, "res:///static")

	p.SetProperty(WebErrorPageContentType, "text/html")
	p.SetProperty(WebErrorPageResource, "res:///static/http404.html")
	p.SetProperty(WebErrorPageStatus, "404")

	p.SetProperty(ServerName, "default")
	p.SetProperty(ServerHost, "0.0.0.0")
	p.SetProperty(ServerPort, "8080")
	p.SetProperty(ServerEnable, "true")

	p.SetProperty(ServerHttpsName, "ssl")
	p.SetProperty(ServerHttpsHost, "0.0.0.0")
	p.SetProperty(ServerHttpsPort, "8443")
	p.SetProperty(ServerHttpsEnable, "false")
	p.SetProperty(ServerHttpsKeyFile, "/ssl/gin-web-server.key")
	p.SetProperty(ServerHttpsCertFile, "/ssl/gin-web-server.crt")

	p.SetProperty(GinMode, "debug")

	return autoGenConfig(cb)
}
