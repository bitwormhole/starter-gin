package glass

import "github.com/bitwormhole/starter/application"

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

	ServerHttpsName   = "server.https.name"
	ServerHttpsPort   = "server.https.port"
	ServerHttpsHost   = "server.https.host"
	ServerHttpsEnable = "server.https.enable"

	GinMode = "gin.mode"

	WebErrorPageResource    = "web.error-page.resource"
	WebErrorPageContentType = "web.error-page.content-type"
	WebErrorPageStatus      = "web.error-page.status"
)

func mainConfig(cb application.ConfigBuilder) error {

	p := cb.DefaultProperties()

	p.SetProperty(WebRestContextPath, "/api")

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

	p.SetProperty(GinMode, "debug")

	return autoGenConfig(cb)
}
