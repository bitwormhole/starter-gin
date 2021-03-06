package glass

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Connector 是服务连接器的接口
type Connector interface {
	Connect(h http.Handler) (NetworkConnection, error)
	Enabled() bool
}

// NetworkConnection 表示 gin.Engine 与网络之间的连接
type NetworkConnection interface {

	// 退出工作循环
	Shutdown() error

	// // 等待工作者循环退出
	// Join() error

	// 运行工作者循环
	Run() error
}

// EngineConnection  表示与gin.Engine 之间的连接
type EngineConnection interface {

	// 取 WebContext
	GetWebContext() WebContext

	//  映射到新的路径
	RequestMapping(path string) EngineConnection

	// 注册过滤器
	Filter(order int, handler gin.HandlerFunc)

	// 注册处理器
	Handle(method string, path string, handler gin.HandlerFunc)

	// 注册 error 处理器
	HandleNoMethod(order int, handler gin.HandlerFunc)

	// 注册 error 处理器
	HandleNoResource(order int, handler gin.HandlerFunc)
}
