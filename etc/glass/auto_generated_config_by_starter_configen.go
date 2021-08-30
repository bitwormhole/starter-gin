// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package glass

import (
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: gin-web-container
	cominfobuilder.Next()
	cominfobuilder.ID("gin-web-container").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theContainer{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: gin-web-server
	cominfobuilder.Next()
	cominfobuilder.ID("gin-web-server").Class("looper").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theServer{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-glass.HTTPConnector
	cominfobuilder.Next()
	cominfobuilder.ID("com2-glass.HTTPConnector").Class("web-server-connector").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theHTTPConnector{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-glass.HTTPSConnector
	cominfobuilder.Next()
	cominfobuilder.ID("com3-glass.HTTPSConnector").Class("web-server-connector").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theHTTPSConnector{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com4-glass.WebContext
	cominfobuilder.Next()
	cominfobuilder.ID("com4-glass.WebContext").Class("web-context").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theStaticWebService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com5-glass.WebContext
	cominfobuilder.Next()
	cominfobuilder.ID("com5-glass.WebContext").Class("web-context").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theRestWebService{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com6-glass.StaticWebResourcesController
	cominfobuilder.Next()
	cominfobuilder.ID("com6-glass.StaticWebResourcesController").Class("static-web-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theStaticWebController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: gin-web-content-types
	cominfobuilder.Next()
	cominfobuilder.ID("gin-web-content-types").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theWebContentTypes{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com8-glass.ErrorController
	cominfobuilder.Next()
	cominfobuilder.ID("com8-glass.ErrorController").Class("static-web-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theWebErrorController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com9-glass.ContextBindController
	cominfobuilder.Next()
	cominfobuilder.ID("com9-glass.ContextBindController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4theContextBindController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4theContainer : the factory of component: gin-web-container
type comFactory4theContainer struct {

    mPrototype * glass0x47343f.Container

	
	mAppContextSelector config.InjectionSelector
	mNameSelector config.InjectionSelector
	mServerSelector config.InjectionSelector
	mConnectorsSelector config.InjectionSelector
	mServicesSelector config.InjectionSelector
	mContentTypesSelector config.InjectionSelector
	mIndexPagesSelector config.InjectionSelector

}

func (inst * comFactory4theContainer) init() application.ComponentFactory {

	
	inst.mAppContextSelector = config.NewInjectionSelector("context",nil)
	inst.mNameSelector = config.NewInjectionSelector("${server.name}",nil)
	inst.mServerSelector = config.NewInjectionSelector("#gin-web-server",nil)
	inst.mConnectorsSelector = config.NewInjectionSelector(".web-server-connector",nil)
	inst.mServicesSelector = config.NewInjectionSelector(".web-context",nil)
	inst.mContentTypesSelector = config.NewInjectionSelector("#gin-web-content-types",nil)
	inst.mIndexPagesSelector = config.NewInjectionSelector("${web.static.index-pages}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theContainer) newObject() * glass0x47343f.Container {
	return & glass0x47343f.Container {}
}

func (inst * comFactory4theContainer) castObject(instance application.ComponentInstance) * glass0x47343f.Container {
	return instance.Get().(*glass0x47343f.Container)
}

func (inst * comFactory4theContainer) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theContainer) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theContainer) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theContainer) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theContainer) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theContainer) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AppContext = inst.getterForFieldAppContextSelector(context)
	obj.Name = inst.getterForFieldNameSelector(context)
	obj.Server = inst.getterForFieldServerSelector(context)
	obj.Connectors = inst.getterForFieldConnectorsSelector(context)
	obj.Services = inst.getterForFieldServicesSelector(context)
	obj.ContentTypes = inst.getterForFieldContentTypesSelector(context)
	obj.IndexPages = inst.getterForFieldIndexPagesSelector(context)
	return context.LastError()
}

//getterForFieldAppContextSelector
func (inst * comFactory4theContainer) getterForFieldAppContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldNameSelector
func (inst * comFactory4theContainer) getterForFieldNameSelector (context application.InstanceContext) string {
    return inst.mNameSelector.GetString(context)
}

//getterForFieldServerSelector
func (inst * comFactory4theContainer) getterForFieldServerSelector (context application.InstanceContext) *glass0x47343f.Server {

	o1 := inst.mServerSelector.GetOne(context)
	o2, ok := o1.(*glass0x47343f.Server)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "gin-web-container")
		eb.Set("field", "Server")
		eb.Set("type1", "?")
		eb.Set("type2", "*glass0x47343f.Server")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldConnectorsSelector
func (inst * comFactory4theContainer) getterForFieldConnectorsSelector (context application.InstanceContext) []glass0x47343f.Connector {
	list1 := inst.mConnectorsSelector.GetList(context)
	list2 := make([]glass0x47343f.Connector, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(glass0x47343f.Connector)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldServicesSelector
func (inst * comFactory4theContainer) getterForFieldServicesSelector (context application.InstanceContext) []*glass0x47343f.WebContext {
	list1 := inst.mServicesSelector.GetList(context)
	list2 := make([]*glass0x47343f.WebContext, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(*glass0x47343f.WebContext)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldContentTypesSelector
func (inst * comFactory4theContainer) getterForFieldContentTypesSelector (context application.InstanceContext) glass0x47343f.ContentTypeManager {

	o1 := inst.mContentTypesSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.ContentTypeManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "gin-web-container")
		eb.Set("field", "ContentTypes")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.ContentTypeManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldIndexPagesSelector
func (inst * comFactory4theContainer) getterForFieldIndexPagesSelector (context application.InstanceContext) string {
    return inst.mIndexPagesSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theServer : the factory of component: gin-web-server
type comFactory4theServer struct {

    mPrototype * glass0x47343f.Server

	
	mContainerSelector config.InjectionSelector

}

func (inst * comFactory4theServer) init() application.ComponentFactory {

	
	inst.mContainerSelector = config.NewInjectionSelector("#gin-web-container",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theServer) newObject() * glass0x47343f.Server {
	return & glass0x47343f.Server {}
}

func (inst * comFactory4theServer) castObject(instance application.ComponentInstance) * glass0x47343f.Server {
	return instance.Get().(*glass0x47343f.Server)
}

func (inst * comFactory4theServer) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theServer) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theServer) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theServer) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Start()
}

func (inst * comFactory4theServer) Destroy(instance application.ComponentInstance) error {
	return inst.castObject(instance).Stop()
}

func (inst * comFactory4theServer) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Container = inst.getterForFieldContainerSelector(context)
	return context.LastError()
}

//getterForFieldContainerSelector
func (inst * comFactory4theServer) getterForFieldContainerSelector (context application.InstanceContext) *glass0x47343f.Container {

	o1 := inst.mContainerSelector.GetOne(context)
	o2, ok := o1.(*glass0x47343f.Container)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "gin-web-server")
		eb.Set("field", "Container")
		eb.Set("type1", "?")
		eb.Set("type2", "*glass0x47343f.Container")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theHTTPConnector : the factory of component: com2-glass.HTTPConnector
type comFactory4theHTTPConnector struct {

    mPrototype * glass0x47343f.HTTPConnector

	
	mGinModeSelector config.InjectionSelector
	mHostSelector config.InjectionSelector
	mPortSelector config.InjectionSelector
	mEnableSelector config.InjectionSelector

}

func (inst * comFactory4theHTTPConnector) init() application.ComponentFactory {

	
	inst.mGinModeSelector = config.NewInjectionSelector("${gin.mode}",nil)
	inst.mHostSelector = config.NewInjectionSelector("${server.host}",nil)
	inst.mPortSelector = config.NewInjectionSelector("${server.port}",nil)
	inst.mEnableSelector = config.NewInjectionSelector("${server.enable}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theHTTPConnector) newObject() * glass0x47343f.HTTPConnector {
	return & glass0x47343f.HTTPConnector {}
}

func (inst * comFactory4theHTTPConnector) castObject(instance application.ComponentInstance) * glass0x47343f.HTTPConnector {
	return instance.Get().(*glass0x47343f.HTTPConnector)
}

func (inst * comFactory4theHTTPConnector) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theHTTPConnector) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theHTTPConnector) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theHTTPConnector) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theHTTPConnector) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theHTTPConnector) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.GinMode = inst.getterForFieldGinModeSelector(context)
	obj.Host = inst.getterForFieldHostSelector(context)
	obj.Port = inst.getterForFieldPortSelector(context)
	obj.Enable = inst.getterForFieldEnableSelector(context)
	return context.LastError()
}

//getterForFieldGinModeSelector
func (inst * comFactory4theHTTPConnector) getterForFieldGinModeSelector (context application.InstanceContext) string {
    return inst.mGinModeSelector.GetString(context)
}

//getterForFieldHostSelector
func (inst * comFactory4theHTTPConnector) getterForFieldHostSelector (context application.InstanceContext) string {
    return inst.mHostSelector.GetString(context)
}

//getterForFieldPortSelector
func (inst * comFactory4theHTTPConnector) getterForFieldPortSelector (context application.InstanceContext) int {
    return inst.mPortSelector.GetInt(context)
}

//getterForFieldEnableSelector
func (inst * comFactory4theHTTPConnector) getterForFieldEnableSelector (context application.InstanceContext) bool {
    return inst.mEnableSelector.GetBool(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theHTTPSConnector : the factory of component: com3-glass.HTTPSConnector
type comFactory4theHTTPSConnector struct {

    mPrototype * glass0x47343f.HTTPSConnector

	
	mGinModeSelector config.InjectionSelector
	mHostSelector config.InjectionSelector
	mPortSelector config.InjectionSelector
	mEnableSelector config.InjectionSelector
	mCertificateFileSelector config.InjectionSelector
	mPrivateKeyFileSelector config.InjectionSelector

}

func (inst * comFactory4theHTTPSConnector) init() application.ComponentFactory {

	
	inst.mGinModeSelector = config.NewInjectionSelector("${gin.mode}",nil)
	inst.mHostSelector = config.NewInjectionSelector("${server.https.host}",nil)
	inst.mPortSelector = config.NewInjectionSelector("${server.https.port}",nil)
	inst.mEnableSelector = config.NewInjectionSelector("${server.https.enable}",nil)
	inst.mCertificateFileSelector = config.NewInjectionSelector("${server.https.cert-file}",nil)
	inst.mPrivateKeyFileSelector = config.NewInjectionSelector("${server.https.key-file}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theHTTPSConnector) newObject() * glass0x47343f.HTTPSConnector {
	return & glass0x47343f.HTTPSConnector {}
}

func (inst * comFactory4theHTTPSConnector) castObject(instance application.ComponentInstance) * glass0x47343f.HTTPSConnector {
	return instance.Get().(*glass0x47343f.HTTPSConnector)
}

func (inst * comFactory4theHTTPSConnector) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theHTTPSConnector) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theHTTPSConnector) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theHTTPSConnector) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theHTTPSConnector) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theHTTPSConnector) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.GinMode = inst.getterForFieldGinModeSelector(context)
	obj.Host = inst.getterForFieldHostSelector(context)
	obj.Port = inst.getterForFieldPortSelector(context)
	obj.Enable = inst.getterForFieldEnableSelector(context)
	obj.CertificateFile = inst.getterForFieldCertificateFileSelector(context)
	obj.PrivateKeyFile = inst.getterForFieldPrivateKeyFileSelector(context)
	return context.LastError()
}

//getterForFieldGinModeSelector
func (inst * comFactory4theHTTPSConnector) getterForFieldGinModeSelector (context application.InstanceContext) string {
    return inst.mGinModeSelector.GetString(context)
}

//getterForFieldHostSelector
func (inst * comFactory4theHTTPSConnector) getterForFieldHostSelector (context application.InstanceContext) string {
    return inst.mHostSelector.GetString(context)
}

//getterForFieldPortSelector
func (inst * comFactory4theHTTPSConnector) getterForFieldPortSelector (context application.InstanceContext) int {
    return inst.mPortSelector.GetInt(context)
}

//getterForFieldEnableSelector
func (inst * comFactory4theHTTPSConnector) getterForFieldEnableSelector (context application.InstanceContext) bool {
    return inst.mEnableSelector.GetBool(context)
}

//getterForFieldCertificateFileSelector
func (inst * comFactory4theHTTPSConnector) getterForFieldCertificateFileSelector (context application.InstanceContext) string {
    return inst.mCertificateFileSelector.GetString(context)
}

//getterForFieldPrivateKeyFileSelector
func (inst * comFactory4theHTTPSConnector) getterForFieldPrivateKeyFileSelector (context application.InstanceContext) string {
    return inst.mPrivateKeyFileSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theStaticWebService : the factory of component: com4-glass.WebContext
type comFactory4theStaticWebService struct {

    mPrototype * glass0x47343f.WebContext

	
	mContainerSelector config.InjectionSelector
	mContextPathSelector config.InjectionSelector
	mControllersSelector config.InjectionSelector

}

func (inst * comFactory4theStaticWebService) init() application.ComponentFactory {

	
	inst.mContainerSelector = config.NewInjectionSelector("#gin-web-container",nil)
	inst.mContextPathSelector = config.NewInjectionSelector("${web.static.context-path}",nil)
	inst.mControllersSelector = config.NewInjectionSelector(".static-web-controller",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theStaticWebService) newObject() * glass0x47343f.WebContext {
	return & glass0x47343f.WebContext {}
}

func (inst * comFactory4theStaticWebService) castObject(instance application.ComponentInstance) * glass0x47343f.WebContext {
	return instance.Get().(*glass0x47343f.WebContext)
}

func (inst * comFactory4theStaticWebService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theStaticWebService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theStaticWebService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theStaticWebService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theStaticWebService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theStaticWebService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Container = inst.getterForFieldContainerSelector(context)
	obj.ContextPath = inst.getterForFieldContextPathSelector(context)
	obj.Controllers = inst.getterForFieldControllersSelector(context)
	return context.LastError()
}

//getterForFieldContainerSelector
func (inst * comFactory4theStaticWebService) getterForFieldContainerSelector (context application.InstanceContext) *glass0x47343f.Container {

	o1 := inst.mContainerSelector.GetOne(context)
	o2, ok := o1.(*glass0x47343f.Container)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com4-glass.WebContext")
		eb.Set("field", "Container")
		eb.Set("type1", "?")
		eb.Set("type2", "*glass0x47343f.Container")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldContextPathSelector
func (inst * comFactory4theStaticWebService) getterForFieldContextPathSelector (context application.InstanceContext) string {
    return inst.mContextPathSelector.GetString(context)
}

//getterForFieldControllersSelector
func (inst * comFactory4theStaticWebService) getterForFieldControllersSelector (context application.InstanceContext) []glass0x47343f.Controller {
	list1 := inst.mControllersSelector.GetList(context)
	list2 := make([]glass0x47343f.Controller, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(glass0x47343f.Controller)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theRestWebService : the factory of component: com5-glass.WebContext
type comFactory4theRestWebService struct {

    mPrototype * glass0x47343f.WebContext

	
	mContainerSelector config.InjectionSelector
	mContextPathSelector config.InjectionSelector
	mControllersSelector config.InjectionSelector

}

func (inst * comFactory4theRestWebService) init() application.ComponentFactory {

	
	inst.mContainerSelector = config.NewInjectionSelector("#gin-web-container",nil)
	inst.mContextPathSelector = config.NewInjectionSelector("${web.rest.context-path}",nil)
	inst.mControllersSelector = config.NewInjectionSelector(".rest-controller",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theRestWebService) newObject() * glass0x47343f.WebContext {
	return & glass0x47343f.WebContext {}
}

func (inst * comFactory4theRestWebService) castObject(instance application.ComponentInstance) * glass0x47343f.WebContext {
	return instance.Get().(*glass0x47343f.WebContext)
}

func (inst * comFactory4theRestWebService) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theRestWebService) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theRestWebService) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theRestWebService) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theRestWebService) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theRestWebService) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Container = inst.getterForFieldContainerSelector(context)
	obj.ContextPath = inst.getterForFieldContextPathSelector(context)
	obj.Controllers = inst.getterForFieldControllersSelector(context)
	return context.LastError()
}

//getterForFieldContainerSelector
func (inst * comFactory4theRestWebService) getterForFieldContainerSelector (context application.InstanceContext) *glass0x47343f.Container {

	o1 := inst.mContainerSelector.GetOne(context)
	o2, ok := o1.(*glass0x47343f.Container)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com5-glass.WebContext")
		eb.Set("field", "Container")
		eb.Set("type1", "?")
		eb.Set("type2", "*glass0x47343f.Container")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldContextPathSelector
func (inst * comFactory4theRestWebService) getterForFieldContextPathSelector (context application.InstanceContext) string {
    return inst.mContextPathSelector.GetString(context)
}

//getterForFieldControllersSelector
func (inst * comFactory4theRestWebService) getterForFieldControllersSelector (context application.InstanceContext) []glass0x47343f.Controller {
	list1 := inst.mControllersSelector.GetList(context)
	list2 := make([]glass0x47343f.Controller, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(glass0x47343f.Controller)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theStaticWebController : the factory of component: com6-glass.StaticWebResourcesController
type comFactory4theStaticWebController struct {

    mPrototype * glass0x47343f.StaticWebResourcesController

	
	mRootSelector config.InjectionSelector
	mContainerSelector config.InjectionSelector

}

func (inst * comFactory4theStaticWebController) init() application.ComponentFactory {

	
	inst.mRootSelector = config.NewInjectionSelector("${web.static.root}",nil)
	inst.mContainerSelector = config.NewInjectionSelector("#gin-web-container",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theStaticWebController) newObject() * glass0x47343f.StaticWebResourcesController {
	return & glass0x47343f.StaticWebResourcesController {}
}

func (inst * comFactory4theStaticWebController) castObject(instance application.ComponentInstance) * glass0x47343f.StaticWebResourcesController {
	return instance.Get().(*glass0x47343f.StaticWebResourcesController)
}

func (inst * comFactory4theStaticWebController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theStaticWebController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theStaticWebController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theStaticWebController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theStaticWebController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theStaticWebController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Root = inst.getterForFieldRootSelector(context)
	obj.Container = inst.getterForFieldContainerSelector(context)
	return context.LastError()
}

//getterForFieldRootSelector
func (inst * comFactory4theStaticWebController) getterForFieldRootSelector (context application.InstanceContext) string {
    return inst.mRootSelector.GetString(context)
}

//getterForFieldContainerSelector
func (inst * comFactory4theStaticWebController) getterForFieldContainerSelector (context application.InstanceContext) *glass0x47343f.Container {

	o1 := inst.mContainerSelector.GetOne(context)
	o2, ok := o1.(*glass0x47343f.Container)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com6-glass.StaticWebResourcesController")
		eb.Set("field", "Container")
		eb.Set("type1", "?")
		eb.Set("type2", "*glass0x47343f.Container")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theWebContentTypes : the factory of component: gin-web-content-types
type comFactory4theWebContentTypes struct {

    mPrototype * glass0x47343f.DefaultContentTypeManager

	
	mAppContextSelector config.InjectionSelector
	mTypesPropertiesSelector config.InjectionSelector

}

func (inst * comFactory4theWebContentTypes) init() application.ComponentFactory {

	
	inst.mAppContextSelector = config.NewInjectionSelector("context",nil)
	inst.mTypesPropertiesSelector = config.NewInjectionSelector("${web.static.content-types-properties}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theWebContentTypes) newObject() * glass0x47343f.DefaultContentTypeManager {
	return & glass0x47343f.DefaultContentTypeManager {}
}

func (inst * comFactory4theWebContentTypes) castObject(instance application.ComponentInstance) * glass0x47343f.DefaultContentTypeManager {
	return instance.Get().(*glass0x47343f.DefaultContentTypeManager)
}

func (inst * comFactory4theWebContentTypes) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theWebContentTypes) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theWebContentTypes) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theWebContentTypes) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theWebContentTypes) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theWebContentTypes) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AppContext = inst.getterForFieldAppContextSelector(context)
	obj.TypesProperties = inst.getterForFieldTypesPropertiesSelector(context)
	return context.LastError()
}

//getterForFieldAppContextSelector
func (inst * comFactory4theWebContentTypes) getterForFieldAppContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldTypesPropertiesSelector
func (inst * comFactory4theWebContentTypes) getterForFieldTypesPropertiesSelector (context application.InstanceContext) string {
    return inst.mTypesPropertiesSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theWebErrorController : the factory of component: com8-glass.ErrorController
type comFactory4theWebErrorController struct {

    mPrototype * glass0x47343f.ErrorController

	
	mResourcePathSelector config.InjectionSelector
	mContentTypeSelector config.InjectionSelector
	mStatusSelector config.InjectionSelector
	mContainerSelector config.InjectionSelector

}

func (inst * comFactory4theWebErrorController) init() application.ComponentFactory {

	
	inst.mResourcePathSelector = config.NewInjectionSelector("${web.error-page.resource}",nil)
	inst.mContentTypeSelector = config.NewInjectionSelector("${web.error-page.content-type}",nil)
	inst.mStatusSelector = config.NewInjectionSelector("${web.error-page.status}",nil)
	inst.mContainerSelector = config.NewInjectionSelector("#gin-web-container",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theWebErrorController) newObject() * glass0x47343f.ErrorController {
	return & glass0x47343f.ErrorController {}
}

func (inst * comFactory4theWebErrorController) castObject(instance application.ComponentInstance) * glass0x47343f.ErrorController {
	return instance.Get().(*glass0x47343f.ErrorController)
}

func (inst * comFactory4theWebErrorController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theWebErrorController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theWebErrorController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theWebErrorController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theWebErrorController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theWebErrorController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ResourcePath = inst.getterForFieldResourcePathSelector(context)
	obj.ContentType = inst.getterForFieldContentTypeSelector(context)
	obj.Status = inst.getterForFieldStatusSelector(context)
	obj.Container = inst.getterForFieldContainerSelector(context)
	return context.LastError()
}

//getterForFieldResourcePathSelector
func (inst * comFactory4theWebErrorController) getterForFieldResourcePathSelector (context application.InstanceContext) string {
    return inst.mResourcePathSelector.GetString(context)
}

//getterForFieldContentTypeSelector
func (inst * comFactory4theWebErrorController) getterForFieldContentTypeSelector (context application.InstanceContext) string {
    return inst.mContentTypeSelector.GetString(context)
}

//getterForFieldStatusSelector
func (inst * comFactory4theWebErrorController) getterForFieldStatusSelector (context application.InstanceContext) int {
    return inst.mStatusSelector.GetInt(context)
}

//getterForFieldContainerSelector
func (inst * comFactory4theWebErrorController) getterForFieldContainerSelector (context application.InstanceContext) *glass0x47343f.Container {

	o1 := inst.mContainerSelector.GetOne(context)
	o2, ok := o1.(*glass0x47343f.Container)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com8-glass.ErrorController")
		eb.Set("field", "Container")
		eb.Set("type1", "?")
		eb.Set("type2", "*glass0x47343f.Container")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4theContextBindController : the factory of component: com9-glass.ContextBindController
type comFactory4theContextBindController struct {

    mPrototype * glass0x47343f.ContextBindController

	
	mAppContextSelector config.InjectionSelector

}

func (inst * comFactory4theContextBindController) init() application.ComponentFactory {

	
	inst.mAppContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4theContextBindController) newObject() * glass0x47343f.ContextBindController {
	return & glass0x47343f.ContextBindController {}
}

func (inst * comFactory4theContextBindController) castObject(instance application.ComponentInstance) * glass0x47343f.ContextBindController {
	return instance.Get().(*glass0x47343f.ContextBindController)
}

func (inst * comFactory4theContextBindController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4theContextBindController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4theContextBindController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4theContextBindController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theContextBindController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4theContextBindController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AppContext = inst.getterForFieldAppContextSelector(context)
	return context.LastError()
}

//getterForFieldAppContextSelector
func (inst * comFactory4theContextBindController) getterForFieldAppContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}




