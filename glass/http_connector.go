package glass

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// HTTPConnector 是一个基本的http连接器，提供不加密的http服务
type HTTPConnector struct {
	GinMode string

	Host   string
	Port   int
	Enable bool
}

func (inst *HTTPConnector) _Impl() Connector {
	return inst
}

// Open 打开连接
func (inst *HTTPConnector) Open() (EngineConnection, error) {
	agent := &ginEngineAgent{}
	agent.SetMode(inst.GinMode)
	agent.runner = func(engine *gin.Engine) error { return inst.run(engine) }
	if inst.Enable {
		return agent.open()
	}
	return agent.openMock()
}

func (inst *HTTPConnector) run(e *gin.Engine) error {
	host := inst.Host
	port := inst.Port
	addr := host + ":" + strconv.Itoa(port)
	return e.Run(addr)
}
