package gin

import (
	"github.com/bitwormhole/starter-gin/web"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

func Config(cb application.ConfigBuilder) error {

	dp := cb.DefaultProperties()

	dp.SetProperty("server.not-found-status-code", "404")
	dp.SetProperty("server.host", "0.0.0.0")
	dp.SetProperty("server.port", "8080")
	dp.SetProperty("server.context-path", "/")
	dp.SetProperty("server.error-page-name", "error.html")
	dp.SetProperty("server.index-page-names", "index.html,index.htm")
	dp.SetProperty("server.content-type-properties", "resources:///content-type.properties")

	return autoGenConfig(cb)
}

type ginServer struct {
	markup.Component
	instance *web.GinStarter `id:"gin-web-server" class:"looper" initMethod:"Start" destroyMethod:"Stop" `

	NotFoundStatusCode    int                 `inject:"${server.not-found-status-code}"`
	Port                  int                 `inject:"${server.port}"`
	Host                  string              `inject:"${server.host}"`
	ContextPath           string              `inject:"${server.context-path}"`
	ErrorPageName         string              `inject:"${server.error-page-name}"`
	IndexPageNames        string              `inject:"${server.index-page-names}"`
	ContentTypeProperties string              `inject:"${server.content-type-properties}"`
	ApplicationContext    application.Context `inject:"context"`
}
