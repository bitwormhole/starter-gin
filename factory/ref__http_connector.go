package factory

// import (
// 	"context"
// 	"crypto/tls"
// 	"errors"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"strconv"
// 	"syscall"
// 	"time"

// 	"github.com/bitwormhole/starter/vlog"
// 	"github.com/gin-gonic/gin"
// )

// ////////////////////////////////////////////////////////////////////////////////

// // HTTPSConnector 是一个加密的https连接器，提供安全的https服务(如果没有配置证书，则运行不加密的http)
// type HTTPSConnector struct {
// 	Host   string
// 	Port   int
// 	Enable bool

// 	GinMode string

// 	CertificateFile string
// 	PrivateKeyFile  string
// }

// func (inst *HTTPSConnector) _Impl() Connector {
// 	return inst
// }

// // Open 打开连接
// func (inst *HTTPSConnector) Open() (EngineConnection, error) {
// 	agent := &ginEngineAgent{}
// 	agent.SetMode(inst.GinMode)
// 	agent.runner = func(engine *gin.Engine) error { return inst.run(engine) }
// 	if inst.Enable {
// 		return agent.open()
// 	}
// 	return agent.openMock()
// }

// func (inst *HTTPSConnector) run(e *gin.Engine) error {
// 	host := inst.Host
// 	port := inst.Port
// 	addr := host + ":" + strconv.Itoa(port)
// 	tls := inst.loadTLSConfig()
// 	return inst.run2(e, addr, tls)
// }

// func (inst *HTTPSConnector) loadTLSConfig() *tls.Config {
// 	cert := inst.CertificateFile
// 	key := inst.PrivateKeyFile
// 	if cert == key && key == "" {
// 		return nil
// 	}
// 	certificate, err := tls.LoadX509KeyPair(cert, key)
// 	if err != nil {
// 		vlog.Error(err)
// 		return nil
// 	}
// 	config := &tls.Config{Certificates: []tls.Certificate{certificate}}
// 	return config
// }

// func (inst *HTTPSConnector) run2(router *gin.Engine, addr string, tls *tls.Config) error {

// 	if tls == nil {
// 		vlog.Info("Run HTTP@", addr)
// 	} else {
// 		vlog.Info("Run HTTPS@", addr)
// 	}

// 	srv := &http.Server{
// 		Addr:      addr,
// 		Handler:   router,
// 		TLSConfig: tls,
// 	}

// 	// Initializing the server in a goroutine so that
// 	// it won't block the graceful shutdown handling below
// 	go func() {
// 		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
// 			log.Printf("listen: %s\n", err)
// 		}
// 	}()

// 	// Wait for interrupt signal to gracefully shutdown the server with
// 	// a timeout of 5 seconds.
// 	quit := make(chan os.Signal)
// 	// kill (no param) default send syscall.SIGTERM
// 	// kill -2 is syscall.SIGINT
// 	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 	<-quit
// 	log.Println("Shutting down server...")

// 	// The context is used to inform the server it has 5 seconds to finish
// 	// the request it is currently handling
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	if err := srv.Shutdown(ctx); err != nil {
// 		log.Fatal("Server forced to shutdown:", err)
// 	}

// 	log.Println("Server exiting")
// 	return nil
// }

// ////////////////////////////////////////////////////////////////////////////////
