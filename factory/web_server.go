package factory

import (
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

////////////////////////////////////////////////////////////////////////////////

type DefaultServer struct {
	markup.Component `id:"gin-web-server" class:"looper" initMethod:"Start" destroyMethod:"Stop"`

	runtimeHolder DefaultServerRuntimeHolder

	Container glass.Container `inject:"#gin-web-container"`
}

func (inst *DefaultServer) _Impl() (glass.Server, application.Looper) {
	return inst, inst
}

func (inst *DefaultServer) GetContainer() glass.Container {
	return inst.Container
}

func (inst *DefaultServer) Start() error {
	return inst.runtimeHolder.start(inst.Container)
}

func (inst *DefaultServer) Stop() error {
	return inst.runtimeHolder.stop()
}

func (inst *DefaultServer) Loop() error {
	return inst.runtimeHolder.loop()
}

////////////////////////////////////////////////////////////////////////////////

type DefaultServerRuntimeHolder struct {
	runtime *webRuntime
}

func (inst *DefaultServerRuntimeHolder) start(con glass.Container) error {
	rt := inst.runtime
	if rt != nil {
		rt.stop()
		rt = nil
	}
	rt = &webRuntime{
		container: con,
	}
	inst.runtime = rt
	return rt.start()
}

func (inst *DefaultServerRuntimeHolder) stop() error {
	rt := inst.runtime
	if rt == nil {
		return nil
	}
	return rt.stop()
}

func (inst *DefaultServerRuntimeHolder) loop() error {
	rt := inst.runtime
	if rt == nil {
		return nil
	}
	return rt.loop()
}

////////////////////////////////////////////////////////////////////////////////
