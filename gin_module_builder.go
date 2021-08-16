package ginstarter

import (
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/lang"
)

// SimpleModuleBuilder 是用于快速创建自定义Gin模块的工具
type SimpleModuleBuilder interface {

	// Name 设置模块的名称
	Name(name string) SimpleModuleBuilder

	// Version 设置模块的版本
	Version(ver string) SimpleModuleBuilder

	// Revision 设置模块的修订号码
	Revision(rev int) SimpleModuleBuilder

	// Resources 设置与模块绑定的资源
	Resources(r collection.Resources) SimpleModuleBuilder

	// Dependency 向模块添加一个依赖
	Dependency(dep application.Module) SimpleModuleBuilder

	// Dependencies 向模块添加一组依赖
	Dependencies(deps []application.Module) SimpleModuleBuilder

	// OnMount 添加一个钩子到 SimpleModuleBuilder
	OnMount(fn application.OnMountFunc) SimpleModuleBuilder

	// RegisterController 注册一个控制器, 【注意】在此注册的控制器无法作为依赖注入的目标，如果要向该控制器注入依赖，请使用标准的模块定义方式！
	RegisterController(ctrl glass.Controller) SimpleModuleBuilder

	// RegisterControllerFunc 以函数形式注册一个控制器
	RegisterControllerFunc(ctrl glass.ControllerFunc) SimpleModuleBuilder

	// Create 创建模块
	Create() application.Module
}

// SimpleModule 函数创建一个 SimpleModuleBuilder，【注意】 此方法无法向控制器注入依赖，如果需要注入，请使用标准的模块定义方式！
func SimpleModule() SimpleModuleBuilder {
	return &simpleGinModuleBuilder{}
}

////////////////////////////////////////////////////////////////////////////////

type simpleGinModuleBuilder struct {
	inner           *application.ModuleBuilder
	group           *glass.GroupController
	onMountHandlers []application.OnMountFunc
}

// Inner 获取内置的 application.ModuleBuilder
func (inst *simpleGinModuleBuilder) Inner() *application.ModuleBuilder {
	inner := inst.inner
	if inner == nil {
		inner = &application.ModuleBuilder{}
		inst.inner = inner
	}
	return inner
}

func (inst *simpleGinModuleBuilder) getGroup() *glass.GroupController {
	g := inst.group
	if g == nil {
		g = &glass.GroupController{}
		inst.group = g
	}
	return g
}

func (inst *simpleGinModuleBuilder) fireOnMountHandlers(cb application.ConfigBuilder) error {
	list := inst.onMountHandlers
	if list == nil {
		return nil
	}
	for _, fn := range list {
		if fn == nil {
			continue
		}
		err := fn(cb)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *simpleGinModuleBuilder) onMount(cb application.ConfigBuilder) error {

	err := inst.fireOnMountHandlers(cb)
	if err != nil {
		return err
	}

	// 注册组件

	groupCtr := inst.getGroup()
	cominfobuilder := &config.ComInfoBuilder{}

	cominfobuilder.Reset()
	cominfobuilder.ID("").Class("rest-controller").Aliases("").Scope("singleton")
	cominfobuilder.OnNew(func() lang.Object { return groupCtr })
	return cominfobuilder.CreateTo(cb)
}

func (inst *simpleGinModuleBuilder) Create() application.Module {
	inner := inst.Inner()
	inner.OnMount(func(cb application.ConfigBuilder) error { return inst.onMount(cb) })
	return inner.Create()
}

func (inst *simpleGinModuleBuilder) Name(name string) SimpleModuleBuilder {
	inst.Inner().Name(name)
	return inst
}

func (inst *simpleGinModuleBuilder) Version(version string) SimpleModuleBuilder {
	inst.Inner().Version(version)
	return inst
}

func (inst *simpleGinModuleBuilder) Revision(revision int) SimpleModuleBuilder {
	inst.Inner().Revision(revision)
	return inst
}

func (inst *simpleGinModuleBuilder) Resources(r collection.Resources) SimpleModuleBuilder {
	inst.Inner().Resources(r)
	return inst
}

func (inst *simpleGinModuleBuilder) OnMount(fn application.OnMountFunc) SimpleModuleBuilder {
	if fn == nil {
		return inst
	}
	list := inst.onMountHandlers
	if list == nil {
		list = make([]application.OnMountFunc, 0)
	}
	list = append(list, fn)
	inst.onMountHandlers = list
	return inst
}

func (inst *simpleGinModuleBuilder) Dependency(dep application.Module) SimpleModuleBuilder {
	inst.Inner().Dependency(dep)
	return inst
}

func (inst *simpleGinModuleBuilder) Dependencies(deps []application.Module) SimpleModuleBuilder {
	inst.Inner().Dependencies(deps)
	return inst
}

func (inst *simpleGinModuleBuilder) RegisterController(ctrl glass.Controller) SimpleModuleBuilder {
	inst.getGroup().AddController(ctrl)
	return inst
}

func (inst *simpleGinModuleBuilder) RegisterControllerFunc(fn glass.ControllerFunc) SimpleModuleBuilder {
	inst.getGroup().AddFunction(fn)
	return inst
}
