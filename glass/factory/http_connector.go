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

	Host     string `inject:"${server.host}"`
	Port     int    `inject:"${server.port}"`
	MyEnable bool   `inject:"${server.enable}"`
}

func (inst *HTTPConnector) _Impl() glass.Connector {
	return inst
}

func (inst *HTTPConnector) Enabled() bool {
	return inst.MyEnable
}

// Connect 打开连接
func (inst *HTTPConnector) Connect(h http.Handler) (glass.NetworkConnection, error) {
	return nil, nil
}
