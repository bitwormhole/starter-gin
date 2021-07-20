package etc_gin

import (
	"github.com/bitwormhole/starter-gin/web"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

type ginServer struct {
	markup.Component
	instance *web.GinStarter `id:"gin-web-server" class:"looper" initMethod:"Start" destroyMethod:"Stop" `

	Host               string              `inject:"${server.host}"`
	Port               int                 `inject:"${server.port}"`
	ContextPath        string              `inject:"${server.context-path}"`
	ApplicationContext application.Context `inject:"context"`
}
