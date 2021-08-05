package web

import "github.com/bitwormhole/starter/collection"

////////////////////////////////////////////////////////////////////////////////

// StaticResourceHolder 表示对某个资源的引用
type StaticResourceHolder interface {
	Error() error
	Ready() bool
	WebURI() string
	ResourceURI() string
	Data() []byte
	Name() string
	ContentType() string
	Size() int
}

////////////////////////////////////////////////////////////////////////////////

// StaticResourceLoader 是一个静态资源加载器
type StaticResourceLoader struct {
	// http-path : '/a/b/c'
	// res-path : 'resources:///path/a/b/c'
	// in-context-path

	// public
	GinServerContext *GinServerContext

	// private
	inner *innerStaticResourceLoader
}

func (inst *StaticResourceLoader) getInnerLoader() *innerStaticResourceLoader {
	inner := inst.inner
	if inner == nil {
		inner = &innerStaticResourceLoader{}
		inner.init(inst.GinServerContext)
		inst.inner = inner
	}
	return inner
}

// Load : 加载资源
func (inst *StaticResourceLoader) Load(uri string) StaticResourceHolder {

	// app := inst.GinServerContext.ApplicationContext
	// res := app.GetResources()

	inner := inst.getInnerLoader()
	holder := &innerResourceHolder{}
	holder.init(inner, uri)
	return holder
}

////////////////////////////////////////////////////////////////////////////////

type innerStaticResourceLoader struct {
	GinServerContext *GinServerContext
	baseWebURI       string
	baseResURI       string // like 'resources:///path/a/b/c'
}

func (inst *innerStaticResourceLoader) init(sc *GinServerContext) {
	cfg := sc.Configuration
	inst.GinServerContext = sc
	inst.baseResURI = cfg.StaticResourceFolder
	inst.baseWebURI = cfg.StaticContextPath
}

////////////////////////////////////////////////////////////////////////////////

type innerResourceHolder struct {
	parent *innerStaticResourceLoader
	res    collection.Resources
	err    error
	data   []byte

	ready       bool
	cacheable   bool
	size        int
	contentType string
	name        string
	resURI      string
	webURI      string
}

func (inst *innerResourceHolder) init(parent *innerStaticResourceLoader, path string) {

	app := parent.GinServerContext.ApplicationContext
	res := app.GetResources()
	baseWebURI := parent.baseWebURI
	baseResURI := parent.baseResURI
	targetWebURI := baseWebURI + path
	targetResURI := baseResURI + path

	data, err := res.GetBinary(targetResURI)
	if err != nil {
		inst.err = err
		inst.ready = false
	} else {

		inst.data = data
		inst.cacheable = true
		inst.size = len(data)
		inst.contentType = ""
		inst.ready = true
	}

	inst.name = ""
	inst.resURI = targetResURI
	inst.webURI = targetWebURI
}

func (inst *innerResourceHolder) Error() error {
	return inst.err
}

func (inst *innerResourceHolder) Ready() bool {
	return inst.ready
}

func (inst *innerResourceHolder) WebURI() string {
	return inst.webURI
}

func (inst *innerResourceHolder) ResourceURI() string {
	return inst.resURI
}

func (inst *innerResourceHolder) Data() []byte {
	return inst.data
}

func (inst *innerResourceHolder) Name() string {
	return inst.name
}

func (inst *innerResourceHolder) ContentType() string {
	return inst.contentType
}

func (inst *innerResourceHolder) Size() int {
	return inst.size
}

////////////////////////////////////////////////////////////////////////////////
