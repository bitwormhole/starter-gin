package glass

import "errors"

// Server 表示 gin-glass 服务器对象
type Server struct {
	Container *Container
}

// Start 启动服务器
func (inst *Server) Start() error {
	rt := inst.Container.runtime
	if rt != nil {
		return errors.New("the container is running")
	}
	rt = &serverRuntime{}
	rt.Container = inst.Container
	inst.Container.runtime = rt
	return rt.start()
}

// Loop 执行服务器主循环
func (inst *Server) Loop() error {
	rt := inst.Container.runtime
	if rt == nil {
		return errors.New("the container is stopped")
	}
	return rt.loop()
}

// Stop 停止服务器
func (inst *Server) Stop() error {
	rt := inst.Container.runtime
	if rt == nil {
		return errors.New("The container is stopped")
	}
	return rt.stop()
}
