package web

import (
	"strings"

	"github.com/bitwormhole/starter/util"
	"github.com/gin-gonic/gin"
)

// StaticResourcesController 是一个静态资源控制器
type StaticResourcesController struct {
	ServerContext *GinServerContext
}

func (inst *StaticResourcesController) _Impl() Controller {
	return inst
}

// Config 方法向container注册控制器本身
func (inst *StaticResourcesController) Config(c Container) error {

	conf := inst.ServerContext.Configuration
	folder := conf.StaticResourceFolder
	all := inst.listHTTPPaths()
	c = c.Mapping(conf.StaticContextPath)

	for index := range all {
		path := all[index]
		holder := &myStaticResHolder{}
		holder.init(inst, path, folder+"/"+path)
		holder.bind(c)
	}

	return nil
}

func (inst *StaticResourcesController) listHTTPPaths() []string {

	conf := inst.ServerContext.Configuration
	app := inst.ServerContext.ApplicationContext
	res := app.GetResources()

	folder := inst.normalizePath(conf.StaticResourceFolder) + "/"
	prefixLength := len(folder)
	all := res.All()
	results := make([]string, 0)

	for index := range all {
		path := all[index]
		path = inst.normalizePath(path)
		if strings.HasPrefix(path, folder) {
			path = path[prefixLength:]
			path = inst.normalizePath(path)
			results = append(results, path)
		}
	}

	return results
}

func (inst *StaticResourcesController) isIndexPage(holder *myStaticResHolder) bool {
	names := inst.ServerContext.indexPageNames
	for index := range names {
		name := names[index]
		suffix := "/" + name
		if strings.HasSuffix(holder.httpPath, suffix) {
			return true
		}
	}
	return false
}

func (inst *StaticResourcesController) normalizePath(path string) string {

	const sep = "://"
	index := strings.Index(path, sep)
	if index > 0 {
		path = path[index+len(sep):]
	}

	builder := &util.PathBuilder{}
	builder.AppendPath(path)
	return builder.String()
}

////////////////////////////////////////////////////////////////////////////////

type myStaticResHolder struct {
	parent   *StaticResourcesController
	source   StaticResourceHolder
	httpPath string
}

func (inst *myStaticResHolder) init(parent *StaticResourcesController, httpPath string, resPath string) {
	loader := parent.ServerContext.ResourceLoader
	inst.source = loader.Load(resPath)
	inst.httpPath = httpPath
	inst.parent = parent
}

func (inst *myStaticResHolder) bind(c Container) {
	if inst.parent == nil {
		return // skip
	}
	path := inst.httpPath
	c.GET(path).Handle(func(c *gin.Context) { inst.handle(c) })
	if inst.parent.isIndexPage(inst) {
		inst.bindAsIndex(c)
	}
}

func (inst *myStaticResHolder) bindAsIndex(c Container) {
	path := inst.httpPath
	i := strings.LastIndex(path, "/")
	path = path[0 : i+1]
	c.GET(path).Handle(func(c *gin.Context) { inst.handle(c) })
}

func (inst *myStaticResHolder) handle(c *gin.Context) {

	src := inst.source
	err := src.Error()

	if err != nil {
		text := err.Error()
		c.JSON(500, text)
		return
	}

	data := src.Data()
	charset := ";charset=utf-8"
	contentType := src.ContentType() + charset
	code := 200
	c.Data(code, contentType, data)
}

////////////////////////////////////////////////////////////////////////////////
