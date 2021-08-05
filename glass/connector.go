package glass

import (
	"io"

	"github.com/gin-gonic/gin"
)

// Connector 是服务连接器的接口
type Connector interface {

	// 打开连接
	Open() (EngineConnection, error)
}

// EngineConnection  表示与gin.Engine 之间的连接
type EngineConnection interface {
	io.Closer // 用于关闭连接

	// 等待工作者循环退出
	Join() error

	// 运行工作者循环
	Run() error

	//  映射到新的路径
	RequestMapping(path string) EngineConnection

	// 注册过滤器
	Filter(order int, handler gin.HandlerFunc)

	// 注册处理器
	Handle(method string, path string, handler gin.HandlerFunc)
}
