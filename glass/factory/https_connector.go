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

	Host     string `inject:"${server.https.host}"`
	Port     int    `inject:"${server.https.port}"`
	MyEnable bool   `inject:"${server.https.enable}"`

	CertificateFile string `inject:"${server.https.cert-file}"`
	PrivateKeyFile  string `inject:"${server.https.key-file}"`
}

func (inst *HTTPSConnector) _Impl() glass.Connector {
	return inst
}

func (inst *HTTPSConnector) Enabled() bool {
	return inst.MyEnable
}

// Connect 打开连接
func (inst *HTTPSConnector) Connect(h http.Handler) (glass.NetworkConnection, error) {
	return nil, nil
}
