// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package etcdemo

import (
	errors "errors"

	elements_095f49 "github.com/bitwormhole/starter-gin/demo/elements"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)

func Config(configbuilder application.ConfigBuilder) error {

	cominfobuilder := &config.ComInfoBuilder{}
	err := errors.New("OK")

	// theGinHandlerExample1
	cominfobuilder.Reset()
	cominfobuilder.ID("theGinHandlerExample1").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &elements_095f49.ExampleGinController{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theGinHandlerExample1{}
		adapter.instance = o.(*elements_095f49.ExampleGinController)
		// adapter.context = context
		err := adapter.__inject__(context)
		if err != nil {
			return err
		}
		return nil
	})
	err = cominfobuilder.CreateTo(configbuilder)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// type theGinHandlerExample1 struct

func (inst *theGinHandlerExample1) __inject__(context application.Context) error {

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

	// to instance

	// invoke custom inject method

	return injection.Close()
}
