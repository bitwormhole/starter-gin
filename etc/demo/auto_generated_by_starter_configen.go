// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package demo

import(
	errors "errors"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)


func autoGenConfig(configbuilder application.ConfigBuilder) error {

	cominfobuilder := &config.ComInfoBuilder{}
	err := errors.New("OK")

    
	// theExampleController1
	cominfobuilder.Reset()
	cominfobuilder.ID("theExampleController1").Class("rest-controller").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &ExampleController1{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theExampleController1{}
		adapter.instance = o.(*ExampleController1)
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

	// theExampleController2
	cominfobuilder.Reset()
	cominfobuilder.ID("theExampleController2").Class("rest-controller").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &ExampleController2{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theExampleController2{}
		adapter.instance = o.(*ExampleController2)
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
// type theExampleController1 struct

func (inst *theExampleController1) __inject__(context application.Context) error {

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

////////////////////////////////////////////////////////////////////////////////
// type theExampleController2 struct

func (inst *theExampleController2) __inject__(context application.Context) error {

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

