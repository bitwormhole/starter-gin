package factory

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
)

// HTTPConnector 是提供HTTP连接器的组件
type HTTPConnector struct {
	markup.Component `class:"web-server-connector"`

	GinMode string `inject:"${gin.mode}"`

	Host   string `inject:"${server.host}"`
	Port   int    `inject:"${server.port}"`
	Enable bool   `inject:"${server.enable}"`
}

func (inst *HTTPConnector) _Impl() glass.Connector {
	return inst
}

// Open 打开连接
func (inst *HTTPConnector) Open(h http.Handler) (glass.NetworkConnection, error) {
	con := &WebConnector{}
	con.Enable = inst.Enable
	con.GinMode = inst.GinMode
	con.Host = inst.Host
	con.Port = inst.Port
	con.CertificateFile = ""
	con.PrivateKeyFile = ""
	con.MainHandler = h
	return con.Open()
}
