// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package devtools

import(
	errors "errors"
	devtools_08d6b0 "github.com/bitwormhole/starter-gin/src/devtools"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)


func autoGenConfig(configbuilder application.ConfigBuilder) error {

	cominfobuilder := &config.ComInfoBuilder{}
	err := errors.New("OK")

    
	// theWebDevtoolsController
	cominfobuilder.Reset()
	cominfobuilder.ID("theWebDevtoolsController").Class("rest-controller").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &devtools_08d6b0.DevtoolsController{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theWebDevtoolsController{}
		adapter.instance = o.(*devtools_08d6b0.DevtoolsController)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
    if err !=nil{
        return err
    }


	return nil
}


////////////////////////////////////////////////////////////////////////////////
// type theWebDevtoolsController struct

func (inst *theWebDevtoolsController) __inject__(context application.Context) error {

	// prepare
	instance := inst.instance
	injection, err := context.Injector().OpenInjection(context)
	if err != nil {
		return err
	}
	defer injection.Close()
	if instance == nil {
		return nil
	}

	// from getters
	inst.AppContext=inst.__get_AppContext__(injection, "context")
	inst.Enable=inst.__get_Enable__(injection, "${web.devtools.enable}")


	// to instance
	instance.AppContext=inst.AppContext
	instance.Enable=inst.Enable


	// invoke custom inject method


	return injection.Close()
}

func (inst * theWebDevtoolsController) __get_AppContext__(injection application.Injection,selector string) application.Context {
	return injection.Context()
}

func (inst * theWebDevtoolsController) __get_Enable__(injection application.Injection,selector string) bool {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadBool()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

