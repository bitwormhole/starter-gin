package glass

// HTTPSConnector 是一个加密的https连接器，提供安全的https服务
type HTTPSConnector struct {
	Host   string
	Port   int
	Enable bool

	CertificateFile string
	PrivateKeyFile  string
}

func (inst *HTTPSConnector) _Impl() Connector {
	return inst
}

// Open 打开连接
func (inst *HTTPSConnector) Open() (EngineConnection, error) {
	agent := &ginEngineAgent{}
	if inst.Enable {
		return agent.open()
	}
	return agent.openMock()
}
