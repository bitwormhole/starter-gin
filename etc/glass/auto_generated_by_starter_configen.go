// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package glass

import(
	errors "errors"
	glass_47343f "github.com/bitwormhole/starter-gin/glass"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)


func autoGenConfig(configbuilder application.ConfigBuilder) error {

	cominfobuilder := &config.ComInfoBuilder{}
	err := errors.New("OK")

    
	// the
	cominfobuilder.Reset()
	cominfobuilder.ID("the").Class("static-web-controller").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &glass_47343f.ErrorController{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &the{}
		adapter.instance = o.(*glass_47343f.ErrorController)
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

	// theContainer
	cominfobuilder.Reset()
	cominfobuilder.ID("gin-web-container").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &glass_47343f.Container{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theContainer{}
		adapter.instance = o.(*glass_47343f.Container)
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

	// theExampleRestController
	cominfobuilder.Reset()
	cominfobuilder.ID("theExampleRestController").Class("rest-controller").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &glass_47343f.DefaultRestController{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theExampleRestController{}
		adapter.instance = o.(*glass_47343f.DefaultRestController)
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

	// theHTTPConnector
	cominfobuilder.Reset()
	cominfobuilder.ID("theHTTPConnector").Class("web-server-connector").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &glass_47343f.HTTPConnector{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theHTTPConnector{}
		adapter.instance = o.(*glass_47343f.HTTPConnector)
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

	// theHTTPSConnector
	cominfobuilder.Reset()
	cominfobuilder.ID("theHTTPSConnector").Class("web-server-connector").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &glass_47343f.HTTPSConnector{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theHTTPSConnector{}
		adapter.instance = o.(*glass_47343f.HTTPSConnector)
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

	// theRestWebService
	cominfobuilder.Reset()
	cominfobuilder.ID("theRestWebService").Class("web-context").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &glass_47343f.WebContext{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theRestWebService{}
		adapter.instance = o.(*glass_47343f.WebContext)
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

	// theServer
	cominfobuilder.Reset()
	cominfobuilder.ID("gin-web-server").Class("looper").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &glass_47343f.Server{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return o.(*glass_47343f.Server).Start()
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return o.(*glass_47343f.Server).Stop()
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theServer{}
		adapter.instance = o.(*glass_47343f.Server)
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

	// theStaticWebController
	cominfobuilder.Reset()
	cominfobuilder.ID("theStaticWebController").Class("static-web-controller").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &glass_47343f.StaticWebResourcesController{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theStaticWebController{}
		adapter.instance = o.(*glass_47343f.StaticWebResourcesController)
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

	// theStaticWebService
	cominfobuilder.Reset()
	cominfobuilder.ID("theStaticWebService").Class("web-context").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &glass_47343f.WebContext{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theStaticWebService{}
		adapter.instance = o.(*glass_47343f.WebContext)
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

	// theWebContentTypes
	cominfobuilder.Reset()
	cominfobuilder.ID("gin-web-content-types").Class("").Scope("").Aliases("")
	cominfobuilder.OnNew(func() lang.Object {
		return &glass_47343f.DefaultContentTypeManager{}
	})
	cominfobuilder.OnInit(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnDestroy(func(o lang.Object) error {
		return nil
	})
	cominfobuilder.OnInject(func(o lang.Object, context application.Context) error {
		adapter := &theWebContentTypes{}
		adapter.instance = o.(*glass_47343f.DefaultContentTypeManager)
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
// type the struct

func (inst *the) __inject__(context application.Context) error {

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
	inst.Container=inst.__get_Container__(injection, "#gin-web-container")
	inst.ContentType=inst.__get_ContentType__(injection, "${web.error-page.content-type}")
	inst.ResourcePath=inst.__get_ResourcePath__(injection, "${web.error-page.resource}")
	inst.Status=inst.__get_Status__(injection, "${web.error-page.status}")


	// to instance
	instance.Container=inst.Container
	instance.ContentType=inst.ContentType
	instance.ResourcePath=inst.ResourcePath
	instance.Status=inst.Status


	// invoke custom inject method


	return injection.Close()
}

func (inst * the) __get_Container__(injection application.Injection,selector string) *glass_47343f.Container {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(*glass_47343f.Container)
	if !ok {
		err := errors.New("cannot cast component instance to type: *glass_47343f.Container")
		injection.OnError(err)
		return nil
	}

	return o2

}

func (inst * the) __get_ContentType__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * the) __get_ResourcePath__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * the) __get_Status__(injection application.Injection,selector string) int {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadInt()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

////////////////////////////////////////////////////////////////////////////////
// type theContainer struct

func (inst *theContainer) __inject__(context application.Context) error {

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
	inst.Connectors=inst.__get_Connectors__(injection, ".web-server-connector")
	inst.ContentTypes=inst.__get_ContentTypes__(injection, "#gin-web-content-types")
	inst.IndexPages=inst.__get_IndexPages__(injection, "${web.static.index-pages}")
	inst.Name=inst.__get_Name__(injection, "${server.name}")
	inst.Server=inst.__get_Server__(injection, "#gin-web-server")
	inst.Services=inst.__get_Services__(injection, ".web-context")


	// to instance
	instance.AppContext=inst.AppContext
	instance.Connectors=inst.Connectors
	instance.ContentTypes=inst.ContentTypes
	instance.IndexPages=inst.IndexPages
	instance.Name=inst.Name
	instance.Server=inst.Server
	instance.Services=inst.Services


	// invoke custom inject method


	return injection.Close()
}

func (inst * theContainer) __get_AppContext__(injection application.Injection,selector string) application.Context {
	return injection.Context()
}

func (inst * theContainer) __get_Connectors__(injection application.Injection,selector string) []glass_47343f.Connector {
	list := make([]glass_47343f.Connector, 0)
	reader := injection.Select(selector)
	defer reader.Close()
	for reader.HasMore() {
		o1, err := reader.Read()
		if err != nil {
			injection.OnError(err)
			return list
		}
		o2, ok := o1.(glass_47343f.Connector)
		if !ok {
			// err = errors.New("bad cast, selector:" + selector)
			// injection.OnError(err)
			// return list
			// warning ...
			continue
		}
		list = append(list, o2)
	}
	return list

}

func (inst * theContainer) __get_ContentTypes__(injection application.Injection,selector string) glass_47343f.ContentTypeManager {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(glass_47343f.ContentTypeManager)
	if !ok {
		err := errors.New("cannot cast component instance to type: glass_47343f.ContentTypeManager")
		injection.OnError(err)
		return nil
	}

	return o2

}

func (inst * theContainer) __get_IndexPages__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * theContainer) __get_Name__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * theContainer) __get_Server__(injection application.Injection,selector string) *glass_47343f.Server {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(*glass_47343f.Server)
	if !ok {
		err := errors.New("cannot cast component instance to type: *glass_47343f.Server")
		injection.OnError(err)
		return nil
	}

	return o2

}

func (inst * theContainer) __get_Services__(injection application.Injection,selector string) []*glass_47343f.WebContext {
	list := make([]*glass_47343f.WebContext, 0)
	reader := injection.Select(selector)
	defer reader.Close()
	for reader.HasMore() {
		o1, err := reader.Read()
		if err != nil {
			injection.OnError(err)
			return list
		}
		o2, ok := o1.(*glass_47343f.WebContext)
		if !ok {
			// err = errors.New("bad cast, selector:" + selector)
			// injection.OnError(err)
			// return list
			// warning ...
			continue
		}
		list = append(list, o2)
	}
	return list

}

////////////////////////////////////////////////////////////////////////////////
// type theExampleRestController struct

func (inst *theExampleRestController) __inject__(context application.Context) error {

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
// type theHTTPConnector struct

func (inst *theHTTPConnector) __inject__(context application.Context) error {

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
	inst.Enable=inst.__get_Enable__(injection, "${server.enable}")
	inst.GinMode=inst.__get_GinMode__(injection, "${gin.mode}")
	inst.Host=inst.__get_Host__(injection, "${server.host}")
	inst.Port=inst.__get_Port__(injection, "${server.port}")


	// to instance
	instance.Enable=inst.Enable
	instance.GinMode=inst.GinMode
	instance.Host=inst.Host
	instance.Port=inst.Port


	// invoke custom inject method


	return injection.Close()
}

func (inst * theHTTPConnector) __get_Enable__(injection application.Injection,selector string) bool {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadBool()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * theHTTPConnector) __get_GinMode__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * theHTTPConnector) __get_Host__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * theHTTPConnector) __get_Port__(injection application.Injection,selector string) int {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadInt()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

////////////////////////////////////////////////////////////////////////////////
// type theHTTPSConnector struct

func (inst *theHTTPSConnector) __inject__(context application.Context) error {

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
	inst.CertificateFile=inst.__get_CertificateFile__(injection, "${server.https.cert-file}")
	inst.Enable=inst.__get_Enable__(injection, "${server.https.enable}")
	inst.GinMode=inst.__get_GinMode__(injection, "${gin.mode}")
	inst.Host=inst.__get_Host__(injection, "${server.https.host}")
	inst.Port=inst.__get_Port__(injection, "${server.https.port}")
	inst.PrivateKeyFile=inst.__get_PrivateKeyFile__(injection, "${server.https.key-file}")


	// to instance
	instance.CertificateFile=inst.CertificateFile
	instance.Enable=inst.Enable
	instance.GinMode=inst.GinMode
	instance.Host=inst.Host
	instance.Port=inst.Port
	instance.PrivateKeyFile=inst.PrivateKeyFile


	// invoke custom inject method


	return injection.Close()
}

func (inst * theHTTPSConnector) __get_CertificateFile__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * theHTTPSConnector) __get_Enable__(injection application.Injection,selector string) bool {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadBool()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * theHTTPSConnector) __get_GinMode__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * theHTTPSConnector) __get_Host__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * theHTTPSConnector) __get_Port__(injection application.Injection,selector string) int {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadInt()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * theHTTPSConnector) __get_PrivateKeyFile__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

////////////////////////////////////////////////////////////////////////////////
// type theRestWebService struct

func (inst *theRestWebService) __inject__(context application.Context) error {

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
	inst.Container=inst.__get_Container__(injection, "#gin-web-container")
	inst.ContextPath=inst.__get_ContextPath__(injection, "${web.rest.context-path}")
	inst.Controllers=inst.__get_Controllers__(injection, ".rest-controller")


	// to instance
	instance.Container=inst.Container
	instance.ContextPath=inst.ContextPath
	instance.Controllers=inst.Controllers


	// invoke custom inject method


	return injection.Close()
}

func (inst * theRestWebService) __get_Container__(injection application.Injection,selector string) *glass_47343f.Container {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(*glass_47343f.Container)
	if !ok {
		err := errors.New("cannot cast component instance to type: *glass_47343f.Container")
		injection.OnError(err)
		return nil
	}

	return o2

}

func (inst * theRestWebService) __get_ContextPath__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * theRestWebService) __get_Controllers__(injection application.Injection,selector string) []glass_47343f.Controller {
	list := make([]glass_47343f.Controller, 0)
	reader := injection.Select(selector)
	defer reader.Close()
	for reader.HasMore() {
		o1, err := reader.Read()
		if err != nil {
			injection.OnError(err)
			return list
		}
		o2, ok := o1.(glass_47343f.Controller)
		if !ok {
			// err = errors.New("bad cast, selector:" + selector)
			// injection.OnError(err)
			// return list
			// warning ...
			continue
		}
		list = append(list, o2)
	}
	return list

}

////////////////////////////////////////////////////////////////////////////////
// type theServer struct

func (inst *theServer) __inject__(context application.Context) error {

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
	inst.Container=inst.__get_Container__(injection, "#gin-web-container")


	// to instance
	instance.Container=inst.Container


	// invoke custom inject method


	return injection.Close()
}

func (inst * theServer) __get_Container__(injection application.Injection,selector string) *glass_47343f.Container {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(*glass_47343f.Container)
	if !ok {
		err := errors.New("cannot cast component instance to type: *glass_47343f.Container")
		injection.OnError(err)
		return nil
	}

	return o2

}

////////////////////////////////////////////////////////////////////////////////
// type theStaticWebController struct

func (inst *theStaticWebController) __inject__(context application.Context) error {

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
	inst.Container=inst.__get_Container__(injection, "#gin-web-container")
	inst.Root=inst.__get_Root__(injection, "${web.static.root}")


	// to instance
	instance.Container=inst.Container
	instance.Root=inst.Root


	// invoke custom inject method


	return injection.Close()
}

func (inst * theStaticWebController) __get_Container__(injection application.Injection,selector string) *glass_47343f.Container {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(*glass_47343f.Container)
	if !ok {
		err := errors.New("cannot cast component instance to type: *glass_47343f.Container")
		injection.OnError(err)
		return nil
	}

	return o2

}

func (inst * theStaticWebController) __get_Root__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

////////////////////////////////////////////////////////////////////////////////
// type theStaticWebService struct

func (inst *theStaticWebService) __inject__(context application.Context) error {

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
	inst.Container=inst.__get_Container__(injection, "#gin-web-container")
	inst.ContextPath=inst.__get_ContextPath__(injection, "${web.static.context-path}")
	inst.Controllers=inst.__get_Controllers__(injection, ".static-web-controller")


	// to instance
	instance.Container=inst.Container
	instance.ContextPath=inst.ContextPath
	instance.Controllers=inst.Controllers


	// invoke custom inject method


	return injection.Close()
}

func (inst * theStaticWebService) __get_Container__(injection application.Injection,selector string) *glass_47343f.Container {

	reader := injection.Select(selector)
	defer reader.Close()

	cnt := reader.Count()
	if cnt != 1 {
		err := errors.New("select.result.count != 1, selector="+selector)
		injection.OnError(err)
		return nil
	}

	o1, err := reader.Read()
	if err != nil {
		injection.OnError(err)
		return nil
	}

	o2, ok := o1.(*glass_47343f.Container)
	if !ok {
		err := errors.New("cannot cast component instance to type: *glass_47343f.Container")
		injection.OnError(err)
		return nil
	}

	return o2

}

func (inst * theStaticWebService) __get_ContextPath__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

func (inst * theStaticWebService) __get_Controllers__(injection application.Injection,selector string) []glass_47343f.Controller {
	list := make([]glass_47343f.Controller, 0)
	reader := injection.Select(selector)
	defer reader.Close()
	for reader.HasMore() {
		o1, err := reader.Read()
		if err != nil {
			injection.OnError(err)
			return list
		}
		o2, ok := o1.(glass_47343f.Controller)
		if !ok {
			// err = errors.New("bad cast, selector:" + selector)
			// injection.OnError(err)
			// return list
			// warning ...
			continue
		}
		list = append(list, o2)
	}
	return list

}

////////////////////////////////////////////////////////////////////////////////
// type theWebContentTypes struct

func (inst *theWebContentTypes) __inject__(context application.Context) error {

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
	inst.TypesProperties=inst.__get_TypesProperties__(injection, "${web.static.content-types-properties}")


	// to instance
	instance.AppContext=inst.AppContext
	instance.TypesProperties=inst.TypesProperties


	// invoke custom inject method


	return injection.Close()
}

func (inst * theWebContentTypes) __get_AppContext__(injection application.Injection,selector string) application.Context {
	return injection.Context()
}

func (inst * theWebContentTypes) __get_TypesProperties__(injection application.Injection,selector string) string {
	reader := injection.Select(selector)
	defer reader.Close()
	value, err := reader.ReadString()
	if err != nil {
		injection.OnError(err)
	}
	return value
}

