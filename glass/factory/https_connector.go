package factory

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
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

// Enabled ...
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
		addr:      addr,
		connector: inst,
		h:         h,
	}
	return nc, nil
}

////////////////////////////////////////////////////////////////////////////////

// HTTPSConn ...
type HTTPSConn struct {
	connector *HTTPSConnector
	h         http.Handler
	addr      string
	stopping  bool
}

func (inst *HTTPSConn) _Impl() glass.NetworkConnection {
	return inst
}

// Shutdown ...
func (inst *HTTPSConn) Shutdown() error {
	inst.stopping = true
	return nil
}

// Run ...
func (inst *HTTPSConn) Run() error {

	vlog.Info("serve HTTPS at [", inst.addr, "]")

	srv := &http.Server{
		Addr:    inst.addr,
		Handler: inst.h,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	cert := inst.connector.CertificateFile
	pkey := inst.connector.PrivateKeyFile
	go func() {
		if err := srv.ListenAndServeTLS(cert, pkey); err != nil && err != http.ErrServerClosed {
			vlog.Error("listen: %s\n", err)
			inst.stopping = true
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	for !inst.stopping {
		time.Sleep(time.Second * 3)
	}
	vlog.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		vlog.Fatal("Server forced to shutdown: ", err)
	}

	vlog.Info("Server exiting")
	return nil
}

////////////////////////////////////////////////////////////////////////////////
