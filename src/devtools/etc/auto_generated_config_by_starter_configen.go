// (todo:gen2.template)
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package etc

import (
	devtools0x08d6b0 "github.com/bitwormhole/starter-gin/src/devtools"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: com0-devtools0x08d6b0.DevtoolsController
	cominfobuilder.Next()
	cominfobuilder.ID("com0-devtools0x08d6b0.DevtoolsController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDevtoolsController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDevtoolsController : the factory of component: com0-devtools0x08d6b0.DevtoolsController
type comFactory4pComDevtoolsController struct {
	mPrototype *devtools0x08d6b0.DevtoolsController

	mAppContextSelector config.InjectionSelector
	mEnableSelector     config.InjectionSelector
}

func (inst *comFactory4pComDevtoolsController) init() application.ComponentFactory {

	inst.mAppContextSelector = config.NewInjectionSelector("context", nil)
	inst.mEnableSelector = config.NewInjectionSelector("${web.devtools.enable}", nil)

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComDevtoolsController) newObject() *devtools0x08d6b0.DevtoolsController {
	return &devtools0x08d6b0.DevtoolsController{}
}

func (inst *comFactory4pComDevtoolsController) castObject(instance application.ComponentInstance) *devtools0x08d6b0.DevtoolsController {
	return instance.Get().(*devtools0x08d6b0.DevtoolsController)
}

func (inst *comFactory4pComDevtoolsController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComDevtoolsController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComDevtoolsController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComDevtoolsController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComDevtoolsController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComDevtoolsController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {

	obj := inst.castObject(instance)
	obj.AppContext = inst.getterForFieldAppContextSelector(context)
	obj.Enable = inst.getterForFieldEnableSelector(context)
	return context.LastError()
}

//getterForFieldAppContextSelector
func (inst *comFactory4pComDevtoolsController) getterForFieldAppContextSelector(context application.InstanceContext) application.Context {
	return context.Context()
}

//getterForFieldEnableSelector
func (inst *comFactory4pComDevtoolsController) getterForFieldEnableSelector(context application.InstanceContext) bool {
	return inst.mEnableSelector.GetBool(context)
}
