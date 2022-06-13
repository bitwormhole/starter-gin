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

// Enabled ...
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

// HTTPConn ...
type HTTPConn struct {
	h        http.Handler
	addr     string
	stopping bool
}

func (inst *HTTPConn) _Impl() glass.NetworkConnection {
	return inst
}

// Shutdown ...
func (inst *HTTPConn) Shutdown() error {
	inst.stopping = true
	return nil
}

// Run ...
func (inst *HTTPConn) Run() error {
	vlog.Info("serve HTTP at [", inst.addr, "]")

	srv := &http.Server{
		Addr:    inst.addr,
		Handler: inst.h,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
