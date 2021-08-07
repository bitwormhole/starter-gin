package glass

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// HTTPSConnector 是一个加密的https连接器，提供安全的https服务
type HTTPSConnector struct {
	Host   string
	Port   int
	Enable bool

	GinMode string

	CertificateFile string
	PrivateKeyFile  string
}

func (inst *HTTPSConnector) _Impl() Connector {
	return inst
}

// Open 打开连接
func (inst *HTTPSConnector) Open() (EngineConnection, error) {
	agent := &ginEngineAgent{}
	agent.SetMode(inst.GinMode)
	agent.runner = func(engine *gin.Engine) error { return inst.run(engine) }
	if inst.Enable {
		return agent.open()
	}
	return agent.openMock()
}

func (inst *HTTPSConnector) run(e *gin.Engine) error {
	host := inst.Host
	port := inst.Port
	cert := inst.CertificateFile
	key := inst.PrivateKeyFile
	addr := host + ":" + strconv.Itoa(port)
	return e.RunTLS(addr, cert, key)
}
