package glass

import "io"

// Connector 是服务连接器的接口
type Connector interface {
	Open() (EngineConnection, error)
}

// EngineConnection  表示与gin.Engine 之间的连接
type EngineConnection interface {
	io.Closer
	RequestMapping(path string) EngineConnection
	Join() error
	Run() error
}
