package glass

import "github.com/bitwormhole/starter/application"

// ExportModule 对外导出本模块
func ExportModule() application.Module {
	return &application.DefineModule{
		Name:    "github.com/bitwormhole/starter-gin/glass",
		Version: "1.0",
		OnMount: func(cb application.ConfigBuilder) error { return mainConfig(cb) },
	}
}

const (
	// WebStaticContextPath  静态web资源的上下文路径
	WebStaticContextPath            = "web.static.context-path"
	WebStaticIndexPages             = "web.static.index-pages"
	WebStaticContentTypesProperties = "web.static.content-types-properties"

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
)

func mainConfig(cb application.ConfigBuilder) error {

	p := cb.DefaultProperties()

	p.SetProperty(WebRestContextPath, "/api")

	p.SetProperty(WebStaticContextPath, "/")
	p.SetProperty(WebStaticIndexPages, "index.html, index.htm")
	p.SetProperty(WebStaticContentTypesProperties, "res:///static-content-types.properties")

	p.SetProperty(ServerName, "www")
	p.SetProperty(ServerHost, "0.0.0.0")
	p.SetProperty(ServerPort, "8080")
	p.SetProperty(ServerEnable, "true")

	p.SetProperty(ServerHttpsName, "ssl")
	p.SetProperty(ServerHttpsHost, "0.0.0.0")
	p.SetProperty(ServerHttpsPort, "8443")
	p.SetProperty(ServerHttpsEnable, "false")

	return autoGenConfig(cb)
}
