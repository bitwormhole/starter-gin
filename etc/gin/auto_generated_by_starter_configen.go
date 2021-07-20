// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package etc_gin

import(
	errors "errors"
	web_9fe371 "github.com/bitwormhole/starter-gin/web"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)


func Config(configbuilder application.ConfigBuilder) error {

	cominfobuilder := &config.ComInfoBuilder{}
	err := errors.New("OK")

    
	// ginServer
	cominfobuilder.Reset()
	cominfobuilder.ID("gin-web-server").Class("looper").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &web_9fe371.GinStarter{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return o.(*web_9fe371.GinStarter).Start()
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return o.(*web_9fe371.GinStarter).Stop()
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &ginServer{}
		adapter.instance = o.(*web_9fe371.GinStarter)
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
// type ginServer struct

func (inst *ginServer) __inject__(context application.Context) error {

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
	inst.ContextPath=inst.__get_ContextPath__(injection, "${server.context-path}")
	inst.Host=inst.__get_Host__(injection, "${server.host}")
	inst.Port=inst.__get_Port__(injection, "${server.port}")


	// to instance
	instance.ApplicationContext=inst.ApplicationContext
	instance.ContextPath=inst.ContextPath
	instance.Host=inst.Host
	instance.Port=inst.Port


	// invoke custom inject method


	return injection.Close()
}

func (inst * ginServer) __get_ApplicationContext__(injection application.Injection,selector string) application.Context {
	return injection.Context()
}

func (inst * ginServer) __get_ContextPath__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * ginServer) __get_Host__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * ginServer) __get_Port__(injection application.Injection,selector string) int {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadInt()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

