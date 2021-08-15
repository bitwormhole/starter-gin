package ginstarter

import (
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/lang"
)

// GinModuleBuilder 是用于快速创建自定义Gin模块的工具，【注意】GinModuleBuilder 无法向控制器注入依赖，如果需要注入，请使用标准的模块定义方式！
type GinModuleBuilder struct {
	inner *application.ModuleBuilder
	group *glass.GroupController
}

// Inner 获取内置的 application.ModuleBuilder
func (inst *GinModuleBuilder) Inner() *application.ModuleBuilder {
	inner := inst.inner
	if inner == nil {
		inner = &application.ModuleBuilder{}
		inst.inner = inner
	}
	return inner
}

func (inst *GinModuleBuilder) getGroup() *glass.GroupController {
	g := inst.group
	if g == nil {
		g = &glass.GroupController{}
		inst.group = g
	}
	return g
}

func (inst *GinModuleBuilder) onMount(cb application.ConfigBuilder) error {

	groupCtr := inst.getGroup()
	cominfobuilder := &config.ComInfoBuilder{}

	cominfobuilder.Reset()
	cominfobuilder.ID("").Class("rest-controller").Aliases("").Scope("singleton")
	cominfobuilder.OnNew(func() lang.Object { return groupCtr })
	return cominfobuilder.CreateTo(cb)
}

// Create 创建模块
func (inst *GinModuleBuilder) Create() application.Module {
	inner := inst.Inner()
	inner.OnMount(func(cb application.ConfigBuilder) error { return inst.onMount(cb) })
	return inner.Create()
}

// Name 设置模块的名称
func (inst *GinModuleBuilder) Name(name string) *GinModuleBuilder {
	inst.Inner().Name(name)
	return inst
}

// Version 设置模块的版本
func (inst *GinModuleBuilder) Version(version string) *GinModuleBuilder {
	inst.Inner().Version(version)
	return inst
}

// Revision 设置模块的修订号码
func (inst *GinModuleBuilder) Revision(revision int) *GinModuleBuilder {
	inst.Inner().Revision(revision)
	return inst
}

// RegisterController 注册一个控制器, 【注意】在此注册的控制器无法作为依赖注入的目标，如果要向该控制器注入依赖，请使用标准的模块定义方式！
func (inst *GinModuleBuilder) RegisterController(ctrl glass.Controller) *GinModuleBuilder {
	inst.getGroup().AddController(ctrl)
	return inst
}

// RegisterControllerFunc 以函数形式注册一个控制器
func (inst *GinModuleBuilder) RegisterControllerFunc(fn glass.ControllerFunc) *GinModuleBuilder {
	inst.getGroup().AddFunction(fn)
	return inst
}
