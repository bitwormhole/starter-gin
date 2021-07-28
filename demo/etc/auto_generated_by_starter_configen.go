// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package etcdemo

import(
	errors "errors"
	elements_095f49 "github.com/bitwormhole/starter-gin/demo/elements"
	web_9fe371 "github.com/bitwormhole/starter-gin/web"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)


func Config(configbuilder application.ConfigBuilder) error {

	cominfobuilder := &config.ComInfoBuilder{}
	err := errors.New("OK")

    
	// the404handler
	cominfobuilder.Reset()
	cominfobuilder.ID("the404handler").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &web_9fe371.Http404pageController{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &the404handler{}
		adapter.instance = o.(*web_9fe371.Http404pageController)
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
    if err !=nil{
        return err
    }

	// theStaticHandler
	cominfobuilder.Reset()
	cominfobuilder.ID("theStaticHandler").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &web_9fe371.StaticResourcesController{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theStaticHandler{}
		adapter.instance = o.(*web_9fe371.StaticResourcesController)
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
// type the404handler struct

func (inst *the404handler) __inject__(context application.Context) error {

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

////////////////////////////////////////////////////////////////////////////////
// type theStaticHandler struct

func (inst *theStaticHandler) __inject__(context application.Context) error {

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
	inst.ApplicationContext=inst.__get_ApplicationContext__(injection, "context")


	// to instance
	instance.ApplicationContext=inst.ApplicationContext


	// invoke custom inject method


	return injection.Close()
}

func (inst * theStaticHandler) __get_ApplicationContext__(injection application.Injection,selector string) application.Context {
	return injection.Context()
}

