// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package config_demo

import (
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	demo0xa69cb5 "github.com/bitwormhole/starter-gin/src/demo/golang/demo"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: com0-demo0xa69cb5.Demo1ctrl
	cominfobuilder.Next()
	cominfobuilder.ID("com0-demo0xa69cb5.Demo1ctrl").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDemo1ctrl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-demo0xa69cb5.Demo1ir
	cominfobuilder.Next()
	cominfobuilder.ID("com1-demo0xa69cb5.Demo1ir").Class("rest-interceptor-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDemo1ir{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDemo1ctrl : the factory of component: com0-demo0xa69cb5.Demo1ctrl
type comFactory4pComDemo1ctrl struct {

    mPrototype * demo0xa69cb5.Demo1ctrl

	
	mRespSelector config.InjectionSelector

}

func (inst * comFactory4pComDemo1ctrl) init() application.ComponentFactory {

	
	inst.mRespSelector = config.NewInjectionSelector("#glass-main-responder",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDemo1ctrl) newObject() * demo0xa69cb5.Demo1ctrl {
	return & demo0xa69cb5.Demo1ctrl {}
}

func (inst * comFactory4pComDemo1ctrl) castObject(instance application.ComponentInstance) * demo0xa69cb5.Demo1ctrl {
	return instance.Get().(*demo0xa69cb5.Demo1ctrl)
}

func (inst * comFactory4pComDemo1ctrl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDemo1ctrl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDemo1ctrl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDemo1ctrl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo1ctrl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo1ctrl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Resp = inst.getterForFieldRespSelector(context)
	return context.LastError()
}

//getterForFieldRespSelector
func (inst * comFactory4pComDemo1ctrl) getterForFieldRespSelector (context application.InstanceContext) glass0x47343f.MainResponder {

	o1 := inst.mRespSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.MainResponder)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com0-demo0xa69cb5.Demo1ctrl")
		eb.Set("field", "Resp")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.MainResponder")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDemo1ir : the factory of component: com1-demo0xa69cb5.Demo1ir
type comFactory4pComDemo1ir struct {

    mPrototype * demo0xa69cb5.Demo1ir

	

}

func (inst * comFactory4pComDemo1ir) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDemo1ir) newObject() * demo0xa69cb5.Demo1ir {
	return & demo0xa69cb5.Demo1ir {}
}

func (inst * comFactory4pComDemo1ir) castObject(instance application.ComponentInstance) * demo0xa69cb5.Demo1ir {
	return instance.Get().(*demo0xa69cb5.Demo1ir)
}

func (inst * comFactory4pComDemo1ir) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDemo1ir) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDemo1ir) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDemo1ir) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo1ir) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDemo1ir) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}




