// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package config_lib

import (
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	factory0x58e669 "github.com/bitwormhole/starter-gin/glass/factory"
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

	// component: gin-web-content-types
	cominfobuilder.Next()
	cominfobuilder.ID("gin-web-content-types").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDefaultContentTypeManager{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-glass0x47343f.ErrorController
	cominfobuilder.Next()
	cominfobuilder.ID("com1-glass0x47343f.ErrorController").Class("static-web-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComErrorController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-factory0x58e669.ContextBindController
	cominfobuilder.Next()
	cominfobuilder.ID("com2-factory0x58e669.ContextBindController").Class("rest-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComContextBindController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-factory0x58e669.HTTPSConnector
	cominfobuilder.Next()
	cominfobuilder.ID("com3-factory0x58e669.HTTPSConnector").Class("web-server-connector").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComHTTPSConnector{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com4-factory0x58e669.HTTPConnector
	cominfobuilder.Next()
	cominfobuilder.ID("com4-factory0x58e669.HTTPConnector").Class("web-server-connector").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComHTTPConnector{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: gin-web-container
	cominfobuilder.Next()
	cominfobuilder.ID("gin-web-container").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComWebContainer{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: rest-web-context
	cominfobuilder.Next()
	cominfobuilder.ID("rest-web-context").Class("web-context").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRestContext{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: static-web-context
	cominfobuilder.Next()
	cominfobuilder.ID("static-web-context").Class("web-context").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComStaticContext{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com8-glass0x47343f.StaticWebResourcesController
	cominfobuilder.Next()
	cominfobuilder.ID("com8-glass0x47343f.StaticWebResourcesController").Class("static-web-controller").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComStaticWebResourcesController{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDefaultContentTypeManager : the factory of component: gin-web-content-types
type comFactory4pComDefaultContentTypeManager struct {

    mPrototype * glass0x47343f.DefaultContentTypeManager

	
	mAppContextSelector config.InjectionSelector
	mTypesPropertiesSelector config.InjectionSelector

}

func (inst * comFactory4pComDefaultContentTypeManager) init() application.ComponentFactory {

	
	inst.mAppContextSelector = config.NewInjectionSelector("context",nil)
	inst.mTypesPropertiesSelector = config.NewInjectionSelector("${web.static.content-types-properties}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDefaultContentTypeManager) newObject() * glass0x47343f.DefaultContentTypeManager {
	return & glass0x47343f.DefaultContentTypeManager {}
}

func (inst * comFactory4pComDefaultContentTypeManager) castObject(instance application.ComponentInstance) * glass0x47343f.DefaultContentTypeManager {
	return instance.Get().(*glass0x47343f.DefaultContentTypeManager)
}

func (inst * comFactory4pComDefaultContentTypeManager) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDefaultContentTypeManager) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDefaultContentTypeManager) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDefaultContentTypeManager) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultContentTypeManager) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDefaultContentTypeManager) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.AppContext = inst.getterForFieldAppContextSelector(context)
	obj.TypesProperties = inst.getterForFieldTypesPropertiesSelector(context)
	return context.LastError()
}

//getterForFieldAppContextSelector
func (inst * comFactory4pComDefaultContentTypeManager) getterForFieldAppContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}

//getterForFieldTypesPropertiesSelector
func (inst * comFactory4pComDefaultContentTypeManager) getterForFieldTypesPropertiesSelector (context application.InstanceContext) string {
    return inst.mTypesPropertiesSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComErrorController : the factory of component: com1-glass0x47343f.ErrorController
type comFactory4pComErrorController struct {

    mPrototype * glass0x47343f.ErrorController

	
	mResourcePathSelector config.InjectionSelector
	mContentTypeSelector config.InjectionSelector
	mStatusSelector config.InjectionSelector
	mContainerSelector config.InjectionSelector
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComErrorController) init() application.ComponentFactory {

	
	inst.mResourcePathSelector = config.NewInjectionSelector("${web.error-page.resource}",nil)
	inst.mContentTypeSelector = config.NewInjectionSelector("${web.error-page.content-type}",nil)
	inst.mStatusSelector = config.NewInjectionSelector("${web.error-page.status}",nil)
	inst.mContainerSelector = config.NewInjectionSelector("#gin-web-container",nil)
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComErrorController) newObject() * glass0x47343f.ErrorController {
	return & glass0x47343f.ErrorController {}
}

func (inst * comFactory4pComErrorController) castObject(instance application.ComponentInstance) * glass0x47343f.ErrorController {
	return instance.Get().(*glass0x47343f.ErrorController)
}

func (inst * comFactory4pComErrorController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComErrorController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComErrorController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComErrorController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComErrorController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComErrorController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ResourcePath = inst.getterForFieldResourcePathSelector(context)
	obj.ContentType = inst.getterForFieldContentTypeSelector(context)
	obj.Status = inst.getterForFieldStatusSelector(context)
	obj.Container = inst.getterForFieldContainerSelector(context)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldResourcePathSelector
func (inst * comFactory4pComErrorController) getterForFieldResourcePathSelector (context application.InstanceContext) string {
    return inst.mResourcePathSelector.GetString(context)
}

//getterForFieldContentTypeSelector
func (inst * comFactory4pComErrorController) getterForFieldContentTypeSelector (context application.InstanceContext) string {
    return inst.mContentTypeSelector.GetString(context)
}

//getterForFieldStatusSelector
func (inst * comFactory4pComErrorController) getterForFieldStatusSelector (context application.InstanceContext) int {
    return inst.mStatusSelector.GetInt(context)
}

//getterForFieldContainerSelector
func (inst * comFactory4pComErrorController) getterForFieldContainerSelector (context application.InstanceContext) glass0x47343f.Container {

	o1 := inst.mContainerSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.Container)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com1-glass0x47343f.ErrorController")
		eb.Set("field", "Container")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.Container")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldContextSelector
func (inst * comFactory4pComErrorController) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComContextBindController : the factory of component: com2-factory0x58e669.ContextBindController
type comFactory4pComContextBindController struct {

    mPrototype * factory0x58e669.ContextBindController

	
	mOrderSelector config.InjectionSelector

}

func (inst * comFactory4pComContextBindController) init() application.ComponentFactory {

	
	inst.mOrderSelector = config.NewInjectionSelector("${webfilter.context.order}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComContextBindController) newObject() * factory0x58e669.ContextBindController {
	return & factory0x58e669.ContextBindController {}
}

func (inst * comFactory4pComContextBindController) castObject(instance application.ComponentInstance) * factory0x58e669.ContextBindController {
	return instance.Get().(*factory0x58e669.ContextBindController)
}

func (inst * comFactory4pComContextBindController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComContextBindController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComContextBindController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComContextBindController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComContextBindController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComContextBindController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Order = inst.getterForFieldOrderSelector(context)
	return context.LastError()
}

//getterForFieldOrderSelector
func (inst * comFactory4pComContextBindController) getterForFieldOrderSelector (context application.InstanceContext) int {
    return inst.mOrderSelector.GetInt(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComHTTPSConnector : the factory of component: com3-factory0x58e669.HTTPSConnector
type comFactory4pComHTTPSConnector struct {

    mPrototype * factory0x58e669.HTTPSConnector

	
	mGinModeSelector config.InjectionSelector
	mHostSelector config.InjectionSelector
	mPortSelector config.InjectionSelector
	mMyEnableSelector config.InjectionSelector
	mCertificateFileSelector config.InjectionSelector
	mPrivateKeyFileSelector config.InjectionSelector

}

func (inst * comFactory4pComHTTPSConnector) init() application.ComponentFactory {

	
	inst.mGinModeSelector = config.NewInjectionSelector("${gin.mode}",nil)
	inst.mHostSelector = config.NewInjectionSelector("${server.https.host}",nil)
	inst.mPortSelector = config.NewInjectionSelector("${server.https.port}",nil)
	inst.mMyEnableSelector = config.NewInjectionSelector("${server.https.enable}",nil)
	inst.mCertificateFileSelector = config.NewInjectionSelector("${server.https.cert-file}",nil)
	inst.mPrivateKeyFileSelector = config.NewInjectionSelector("${server.https.key-file}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComHTTPSConnector) newObject() * factory0x58e669.HTTPSConnector {
	return & factory0x58e669.HTTPSConnector {}
}

func (inst * comFactory4pComHTTPSConnector) castObject(instance application.ComponentInstance) * factory0x58e669.HTTPSConnector {
	return instance.Get().(*factory0x58e669.HTTPSConnector)
}

func (inst * comFactory4pComHTTPSConnector) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComHTTPSConnector) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComHTTPSConnector) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComHTTPSConnector) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTPSConnector) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTPSConnector) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.GinMode = inst.getterForFieldGinModeSelector(context)
	obj.Host = inst.getterForFieldHostSelector(context)
	obj.Port = inst.getterForFieldPortSelector(context)
	obj.MyEnable = inst.getterForFieldMyEnableSelector(context)
	obj.CertificateFile = inst.getterForFieldCertificateFileSelector(context)
	obj.PrivateKeyFile = inst.getterForFieldPrivateKeyFileSelector(context)
	return context.LastError()
}

//getterForFieldGinModeSelector
func (inst * comFactory4pComHTTPSConnector) getterForFieldGinModeSelector (context application.InstanceContext) string {
    return inst.mGinModeSelector.GetString(context)
}

//getterForFieldHostSelector
func (inst * comFactory4pComHTTPSConnector) getterForFieldHostSelector (context application.InstanceContext) string {
    return inst.mHostSelector.GetString(context)
}

//getterForFieldPortSelector
func (inst * comFactory4pComHTTPSConnector) getterForFieldPortSelector (context application.InstanceContext) int {
    return inst.mPortSelector.GetInt(context)
}

//getterForFieldMyEnableSelector
func (inst * comFactory4pComHTTPSConnector) getterForFieldMyEnableSelector (context application.InstanceContext) bool {
    return inst.mMyEnableSelector.GetBool(context)
}

//getterForFieldCertificateFileSelector
func (inst * comFactory4pComHTTPSConnector) getterForFieldCertificateFileSelector (context application.InstanceContext) string {
    return inst.mCertificateFileSelector.GetString(context)
}

//getterForFieldPrivateKeyFileSelector
func (inst * comFactory4pComHTTPSConnector) getterForFieldPrivateKeyFileSelector (context application.InstanceContext) string {
    return inst.mPrivateKeyFileSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComHTTPConnector : the factory of component: com4-factory0x58e669.HTTPConnector
type comFactory4pComHTTPConnector struct {

    mPrototype * factory0x58e669.HTTPConnector

	
	mGinModeSelector config.InjectionSelector
	mHostSelector config.InjectionSelector
	mPortSelector config.InjectionSelector
	mMyEnableSelector config.InjectionSelector

}

func (inst * comFactory4pComHTTPConnector) init() application.ComponentFactory {

	
	inst.mGinModeSelector = config.NewInjectionSelector("${gin.mode}",nil)
	inst.mHostSelector = config.NewInjectionSelector("${server.host}",nil)
	inst.mPortSelector = config.NewInjectionSelector("${server.port}",nil)
	inst.mMyEnableSelector = config.NewInjectionSelector("${server.enable}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComHTTPConnector) newObject() * factory0x58e669.HTTPConnector {
	return & factory0x58e669.HTTPConnector {}
}

func (inst * comFactory4pComHTTPConnector) castObject(instance application.ComponentInstance) * factory0x58e669.HTTPConnector {
	return instance.Get().(*factory0x58e669.HTTPConnector)
}

func (inst * comFactory4pComHTTPConnector) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComHTTPConnector) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComHTTPConnector) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComHTTPConnector) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTPConnector) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTPConnector) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.GinMode = inst.getterForFieldGinModeSelector(context)
	obj.Host = inst.getterForFieldHostSelector(context)
	obj.Port = inst.getterForFieldPortSelector(context)
	obj.MyEnable = inst.getterForFieldMyEnableSelector(context)
	return context.LastError()
}

//getterForFieldGinModeSelector
func (inst * comFactory4pComHTTPConnector) getterForFieldGinModeSelector (context application.InstanceContext) string {
    return inst.mGinModeSelector.GetString(context)
}

//getterForFieldHostSelector
func (inst * comFactory4pComHTTPConnector) getterForFieldHostSelector (context application.InstanceContext) string {
    return inst.mHostSelector.GetString(context)
}

//getterForFieldPortSelector
func (inst * comFactory4pComHTTPConnector) getterForFieldPortSelector (context application.InstanceContext) int {
    return inst.mPortSelector.GetInt(context)
}

//getterForFieldMyEnableSelector
func (inst * comFactory4pComHTTPConnector) getterForFieldMyEnableSelector (context application.InstanceContext) bool {
    return inst.mMyEnableSelector.GetBool(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComWebContainer : the factory of component: gin-web-container
type comFactory4pComWebContainer struct {

    mPrototype * factory0x58e669.WebContainer

	
	mNameSelector config.InjectionSelector
	mIndexPagesSelector config.InjectionSelector
	mContentTypesSelector config.InjectionSelector
	mConnectorsSelector config.InjectionSelector
	mContextsSelector config.InjectionSelector

}

func (inst * comFactory4pComWebContainer) init() application.ComponentFactory {

	
	inst.mNameSelector = config.NewInjectionSelector("${server.name}",nil)
	inst.mIndexPagesSelector = config.NewInjectionSelector("${web.static.index-pages}",nil)
	inst.mContentTypesSelector = config.NewInjectionSelector("#gin-web-content-types",nil)
	inst.mConnectorsSelector = config.NewInjectionSelector(".web-server-connector",nil)
	inst.mContextsSelector = config.NewInjectionSelector(".web-context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComWebContainer) newObject() * factory0x58e669.WebContainer {
	return & factory0x58e669.WebContainer {}
}

func (inst * comFactory4pComWebContainer) castObject(instance application.ComponentInstance) * factory0x58e669.WebContainer {
	return instance.Get().(*factory0x58e669.WebContainer)
}

func (inst * comFactory4pComWebContainer) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComWebContainer) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComWebContainer) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComWebContainer) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWebContainer) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComWebContainer) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Name = inst.getterForFieldNameSelector(context)
	obj.IndexPages = inst.getterForFieldIndexPagesSelector(context)
	obj.ContentTypes = inst.getterForFieldContentTypesSelector(context)
	obj.Connectors = inst.getterForFieldConnectorsSelector(context)
	obj.Contexts = inst.getterForFieldContextsSelector(context)
	return context.LastError()
}

//getterForFieldNameSelector
func (inst * comFactory4pComWebContainer) getterForFieldNameSelector (context application.InstanceContext) string {
    return inst.mNameSelector.GetString(context)
}

//getterForFieldIndexPagesSelector
func (inst * comFactory4pComWebContainer) getterForFieldIndexPagesSelector (context application.InstanceContext) string {
    return inst.mIndexPagesSelector.GetString(context)
}

//getterForFieldContentTypesSelector
func (inst * comFactory4pComWebContainer) getterForFieldContentTypesSelector (context application.InstanceContext) glass0x47343f.ContentTypeManager {

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

//getterForFieldConnectorsSelector
func (inst * comFactory4pComWebContainer) getterForFieldConnectorsSelector (context application.InstanceContext) []glass0x47343f.Connector {
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

//getterForFieldContextsSelector
func (inst * comFactory4pComWebContainer) getterForFieldContextsSelector (context application.InstanceContext) []glass0x47343f.WebContext {
	list1 := inst.mContextsSelector.GetList(context)
	list2 := make([]glass0x47343f.WebContext, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(glass0x47343f.WebContext)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRestContext : the factory of component: rest-web-context
type comFactory4pComRestContext struct {

    mPrototype * factory0x58e669.RestContext

	
	mContextPathSelector config.InjectionSelector
	mContainerSelector config.InjectionSelector
	mControllerListSelector config.InjectionSelector
	mInterceptorsSelector config.InjectionSelector

}

func (inst * comFactory4pComRestContext) init() application.ComponentFactory {

	
	inst.mContextPathSelector = config.NewInjectionSelector("${web.rest.context-path}",nil)
	inst.mContainerSelector = config.NewInjectionSelector("#gin-web-container",nil)
	inst.mControllerListSelector = config.NewInjectionSelector(".rest-controller",nil)
	inst.mInterceptorsSelector = config.NewInjectionSelector(".rest-interceptor-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRestContext) newObject() * factory0x58e669.RestContext {
	return & factory0x58e669.RestContext {}
}

func (inst * comFactory4pComRestContext) castObject(instance application.ComponentInstance) * factory0x58e669.RestContext {
	return instance.Get().(*factory0x58e669.RestContext)
}

func (inst * comFactory4pComRestContext) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRestContext) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRestContext) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRestContext) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRestContext) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRestContext) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ContextPath = inst.getterForFieldContextPathSelector(context)
	obj.Container = inst.getterForFieldContainerSelector(context)
	obj.ControllerList = inst.getterForFieldControllerListSelector(context)
	obj.Interceptors = inst.getterForFieldInterceptorsSelector(context)
	return context.LastError()
}

//getterForFieldContextPathSelector
func (inst * comFactory4pComRestContext) getterForFieldContextPathSelector (context application.InstanceContext) string {
    return inst.mContextPathSelector.GetString(context)
}

//getterForFieldContainerSelector
func (inst * comFactory4pComRestContext) getterForFieldContainerSelector (context application.InstanceContext) glass0x47343f.Container {

	o1 := inst.mContainerSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.Container)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "rest-web-context")
		eb.Set("field", "Container")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.Container")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldControllerListSelector
func (inst * comFactory4pComRestContext) getterForFieldControllerListSelector (context application.InstanceContext) []glass0x47343f.Controller {
	list1 := inst.mControllerListSelector.GetList(context)
	list2 := make([]glass0x47343f.Controller, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(glass0x47343f.Controller)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldInterceptorsSelector
func (inst * comFactory4pComRestContext) getterForFieldInterceptorsSelector (context application.InstanceContext) []glass0x47343f.InterceptorRegistry {
	list1 := inst.mInterceptorsSelector.GetList(context)
	list2 := make([]glass0x47343f.InterceptorRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(glass0x47343f.InterceptorRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComStaticContext : the factory of component: static-web-context
type comFactory4pComStaticContext struct {

    mPrototype * factory0x58e669.StaticContext

	
	mContainerSelector config.InjectionSelector
	mControllerListSelector config.InjectionSelector
	mInterceptorsSelector config.InjectionSelector
	mContextPathSelector config.InjectionSelector

}

func (inst * comFactory4pComStaticContext) init() application.ComponentFactory {

	
	inst.mContainerSelector = config.NewInjectionSelector("#gin-web-container",nil)
	inst.mControllerListSelector = config.NewInjectionSelector(".static-web-controller",nil)
	inst.mInterceptorsSelector = config.NewInjectionSelector(".static-web-interceptor-registry",nil)
	inst.mContextPathSelector = config.NewInjectionSelector("${web.static.context-path}",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComStaticContext) newObject() * factory0x58e669.StaticContext {
	return & factory0x58e669.StaticContext {}
}

func (inst * comFactory4pComStaticContext) castObject(instance application.ComponentInstance) * factory0x58e669.StaticContext {
	return instance.Get().(*factory0x58e669.StaticContext)
}

func (inst * comFactory4pComStaticContext) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComStaticContext) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComStaticContext) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComStaticContext) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComStaticContext) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComStaticContext) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Container = inst.getterForFieldContainerSelector(context)
	obj.ControllerList = inst.getterForFieldControllerListSelector(context)
	obj.Interceptors = inst.getterForFieldInterceptorsSelector(context)
	obj.ContextPath = inst.getterForFieldContextPathSelector(context)
	return context.LastError()
}

//getterForFieldContainerSelector
func (inst * comFactory4pComStaticContext) getterForFieldContainerSelector (context application.InstanceContext) glass0x47343f.Container {

	o1 := inst.mContainerSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.Container)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "static-web-context")
		eb.Set("field", "Container")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.Container")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldControllerListSelector
func (inst * comFactory4pComStaticContext) getterForFieldControllerListSelector (context application.InstanceContext) []glass0x47343f.Controller {
	list1 := inst.mControllerListSelector.GetList(context)
	list2 := make([]glass0x47343f.Controller, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(glass0x47343f.Controller)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldInterceptorsSelector
func (inst * comFactory4pComStaticContext) getterForFieldInterceptorsSelector (context application.InstanceContext) []glass0x47343f.InterceptorRegistry {
	list1 := inst.mInterceptorsSelector.GetList(context)
	list2 := make([]glass0x47343f.InterceptorRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(glass0x47343f.InterceptorRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldContextPathSelector
func (inst * comFactory4pComStaticContext) getterForFieldContextPathSelector (context application.InstanceContext) string {
    return inst.mContextPathSelector.GetString(context)
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComStaticWebResourcesController : the factory of component: com8-glass0x47343f.StaticWebResourcesController
type comFactory4pComStaticWebResourcesController struct {

    mPrototype * glass0x47343f.StaticWebResourcesController

	
	mRootSelector config.InjectionSelector
	mContainerSelector config.InjectionSelector
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComStaticWebResourcesController) init() application.ComponentFactory {

	
	inst.mRootSelector = config.NewInjectionSelector("${web.static.root}",nil)
	inst.mContainerSelector = config.NewInjectionSelector("#gin-web-container",nil)
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComStaticWebResourcesController) newObject() * glass0x47343f.StaticWebResourcesController {
	return & glass0x47343f.StaticWebResourcesController {}
}

func (inst * comFactory4pComStaticWebResourcesController) castObject(instance application.ComponentInstance) * glass0x47343f.StaticWebResourcesController {
	return instance.Get().(*glass0x47343f.StaticWebResourcesController)
}

func (inst * comFactory4pComStaticWebResourcesController) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComStaticWebResourcesController) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComStaticWebResourcesController) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComStaticWebResourcesController) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComStaticWebResourcesController) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComStaticWebResourcesController) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Root = inst.getterForFieldRootSelector(context)
	obj.Container = inst.getterForFieldContainerSelector(context)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldRootSelector
func (inst * comFactory4pComStaticWebResourcesController) getterForFieldRootSelector (context application.InstanceContext) string {
    return inst.mRootSelector.GetString(context)
}

//getterForFieldContainerSelector
func (inst * comFactory4pComStaticWebResourcesController) getterForFieldContainerSelector (context application.InstanceContext) glass0x47343f.Container {

	o1 := inst.mContainerSelector.GetOne(context)
	o2, ok := o1.(glass0x47343f.Container)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com8-glass0x47343f.StaticWebResourcesController")
		eb.Set("field", "Container")
		eb.Set("type1", "?")
		eb.Set("type2", "glass0x47343f.Container")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldContextSelector
func (inst * comFactory4pComStaticWebResourcesController) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}




