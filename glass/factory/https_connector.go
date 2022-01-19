package factory

import (
	"errors"
	"net/http"
	"strconv"

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

	if h == nil {
		return nil, errors.New("http.Handler==nil")
	}
	host := inst.Host
	port := inst.Port
	if port < 1 {
		port = 8443
	}
	addr := host + ":" + strconv.Itoa(port)
	nc := &HTTPSConn{
		h:    h,
		addr: addr,
	}
	return nc, nil
}

////////////////////////////////////////////////////////////////////////////////

type HTTPSConn struct {
	h    http.Handler
	addr string
}

func (inst *HTTPSConn) _Impl() glass.NetworkConnection {
	return inst
}

func (inst *HTTPSConn) Shutdown() error {
	return nil
}

func (inst *HTTPSConn) Run() error {

	// vlog.Info("serve HTTPS at [", inst.addr, "]")
	// return http.ListenAndServe(inst.addr, inst.h)

	return errors.New("HTTPSConnector: no impl")
}

////////////////////////////////////////////////////////////////////////////////
