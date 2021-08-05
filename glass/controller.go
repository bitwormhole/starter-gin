package glass

// Controller 接口表示一个gin-glass控制器
type Controller interface {
	Init(ec EngineConnection) error
}
