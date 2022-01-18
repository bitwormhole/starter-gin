package factory

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
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
	if h == nil {
		return nil, errors.New("http.Handler==nil")
	}
	host := inst.Host
	port := inst.Port
	if port < 1 {
		port = 8080
	}
	addr := host + ":" + strconv.Itoa(port)
	nc := &HTTPConn{
		h:    h,
		addr: addr,
	}
	return nc, nil
}

////////////////////////////////////////////////////////////////////////////////

type HTTPConn struct {
	h    http.Handler
	addr string
}

func (inst *HTTPConn) _Impl() glass.NetworkConnection {
	return inst
}

func (inst *HTTPConn) Shutdown() error {
	return nil
}

func (inst *HTTPConn) Run() error {

	vlog.Info("serve HTTP at [", inst.addr, "]")

	return http.ListenAndServe(inst.addr, inst.h)
}

////////////////////////////////////////////////////////////////////////////////
