package factory

import (
	"crypto/tls"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/bitwormhole/starter-gin/glass"
)

type WebConnector struct {
	GinMode         string
	Host            string
	Port            int
	Enable          bool
	CertificateFile string
	PrivateKeyFile  string
	MainHandler     http.Handler
}

func (inst *WebConnector) connect() (glass.NetworkConnection, error) {
	conn := &webConnection{}
	conn.params = *inst
	return conn.init()
}

////////////////////////////////////////////////////////////////////////////////

type webConnection struct {
	params WebConnector

	server http.Server

	started  bool
	starting bool
	stopped  bool
	stopping bool
}

func (inst *webConnection) _Impl() glass.NetworkConnection {
	return inst
}

func (inst *webConnection) init() (glass.NetworkConnection, error) {

	inst.starting = true

	// for address (host:port)
	host := inst.params.Host
	port := inst.params.Port
	if port < 1 {
		port = 8080
	}
	addr := host + ":" + strconv.Itoa(port)

	// for TLS config (optional)
	cert := inst.params.CertificateFile
	pkey := inst.params.PrivateKeyFile
	if len(cert) > 0 && len(pkey) > 0 {
		keypair, err := tls.LoadX509KeyPair(cert, pkey)
		if err != nil {
			return nil, err
		}
		config := &tls.Config{Certificates: []tls.Certificate{keypair}}
		inst.server.TLSConfig = config
	}

	inst.server.Addr = addr
	inst.server.Handler = inst.params.MainHandler
	return inst, nil
}

func (inst *webConnection) Shutdown() error {
	inst.stopping = true
	return inst.server.Shutdown(nil)
}

func (inst *webConnection) open() error {
	return nil
}

func (inst *webConnection) close() error {
	return nil
}

// 等待工作者循环退出
func (inst *webConnection) Join() error {
	for {
		if !inst.starting {
			return errors.New("this connection is not open")
		}
		if inst.stopped {
			break
		}
		time.Sleep(2000 * time.Millisecond)
	}
	return nil
}

// 运行工作者循环
func (inst *webConnection) Run() error {
	defer func() { inst.stopped = true }()
	err := inst.open()
	if err != nil {
		return err
	}
	defer inst.close()
	inst.started = true
	return inst.server.ListenAndServe()
}
