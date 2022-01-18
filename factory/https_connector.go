package factory

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
)

// HTTPSConnector 是实现 HTTPS 连接器的组件
type HTTPSConnector struct {
	markup.Component `class:"web-server-connector"`

	GinMode string `inject:"${gin.mode}"`

	Host   string `inject:"${server.https.host}"`
	Port   int    `inject:"${server.https.port}"`
	Enable bool   `inject:"${server.https.enable}"`

	CertificateFile string `inject:"${server.https.cert-file}"`
	PrivateKeyFile  string `inject:"${server.https.key-file}"`
}

func (inst *HTTPSConnector) _Impl() glass.Connector {
	return inst
}

// Open 打开连接
func (inst *HTTPSConnector) Open(h http.Handler) (glass.NetworkConnection, error) {
	con := &WebConnector{}
	con.Enable = inst.Enable
	con.GinMode = inst.GinMode
	con.Host = inst.Host
	con.Port = inst.Port
	con.CertificateFile = inst.CertificateFile
	con.PrivateKeyFile = inst.PrivateKeyFile
	con.MainHandler = h
	return con.Open()
}
