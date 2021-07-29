package web

import (
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/util"
	"github.com/gin-gonic/gin"
)

type StaticResourcesController struct {
	ApplicationContext application.Context
	StaticResourcePath string

	typeMapper         map[string]string
	resources          collection.Resources
	defaultContentType string
}

func (inst *StaticResourcesController) __impl__() Controller {
	return inst
}

func (inst *StaticResourcesController) getStringDefault(str string, def string) string {
	if str == "" {
		return def
	}
	return str
}

func (inst *StaticResourcesController) init() {

	app := inst.ApplicationContext
	r := app.GetResources()

	type0 := inst.getStringDefault(inst.defaultContentType, "application/octet-stream")
	basepath := inst.getStringDefault(inst.StaticResourcePath, "/static")

	inst.resources = r
	inst.defaultContentType = type0
	inst.StaticResourcePath = basepath
}

func (inst *StaticResourcesController) Config(c Container) error {

	inst.init()

	r := inst.resources
	all := r.All()

	for index := range all {
		path := all[index]
		holder := &myStaticResHolder{}
		holder.init(inst, inst.StaticResourcePath, path)
		holder.bind(c)
	}

	return nil
}

func (inst *StaticResourcesController) isIndexPage(holder *myStaticResHolder) bool {
	path := holder.pathInStatic
	i1 := strings.LastIndex(path, "/")
	i2 := strings.LastIndex(path, ".")
	if 0 <= i1 && i1 < i2 {
		name := path[i1+1 : i2]
		return name == "index"
	}
	return false
}

func (inst *StaticResourcesController) loadContent(holder *myStaticResHolder) ([]byte, error) {

	res_path := holder.pathInResources
	data, err := inst.resources.GetBinary(res_path)

	if err != nil {
		return nil, err
	}

	filename := inst.getContentFileName(res_path)
	contenttype := inst.queryContentType(filename)

	holder.contentLength = len(data)
	holder.contentType = contenttype
	holder.contentFileName = filename

	return data, nil
}

func (inst *StaticResourcesController) getContentFileName(path string) string {
	index := strings.LastIndex(path, "/")
	if index < 0 {
		return path
	}
	return path[index+1:]
}

func (inst *StaticResourcesController) getContentTypeMapper() map[string]string {
	mapper := inst.typeMapper
	if mapper != nil {
		return mapper
	}
	mapper = make(map[string]string)
	mapper[".txt"] = "text/plain"
	mapper[".htm"] = "text/html"
	mapper[".html"] = "text/html"
	mapper[".ico"] = "image/x-icon"
	mapper[".png"] = "image/png"
	mapper[".jpg"] = "image/jpeg"
	mapper[".gif"] = "image/gif"
	mapper[".css"] = "text/css"
	mapper[".js"] = "application/x-javascript"
	inst.typeMapper = mapper
	return mapper
}

func (inst *StaticResourcesController) queryContentType(name string) string {
	index := strings.LastIndex(name, ".")
	if index < 0 {
		return inst.defaultContentType
	}
	suffix := name[index:]
	suffix = strings.TrimSpace(suffix)
	suffix = strings.ToLower(suffix)
	mapper := inst.getContentTypeMapper()
	t := mapper[suffix]
	if t == "" {
		return inst.defaultContentType
	}
	return t
}

func (inst *StaticResourcesController) normalizePath(path string) string {
	builder := &util.PathBuilder{}
	builder.AppendPath(path)
	return builder.String()
}

////////////////////////////////////////////////////////////////////////////////

type myStaticResHolder struct {
	parent *StaticResourcesController

	pathInResources string
	pathInStatic    string

	content         []byte // cached content
	contentType     string
	contentLength   int
	contentFileName string
}

func (inst *myStaticResHolder) init(parent *StaticResourcesController, path1 string, path2 string) {

	// path1 = basepath
	// path2 = full-res-path

	path1 = parent.normalizePath(path1)
	path2 = parent.normalizePath(path2)
	if !strings.HasPrefix(path2, path1) {
		return // skip
	}
	webpath := path2[len(path1):]

	inst.pathInResources = path2
	inst.pathInStatic = webpath
	inst.parent = parent
}

func (inst *myStaticResHolder) bind(c Container) {
	if inst.parent == nil {
		return // skip
	}
	path := "/" + inst.pathInStatic
	c.GET(path).Handle(func(c *gin.Context) { inst.handle(c) })
	if inst.parent.isIndexPage(inst) {
		inst.bindAsIndex(c)
	}
}

func (inst *myStaticResHolder) bindAsIndex(c Container) {
	path := "/" + inst.pathInStatic
	i := strings.LastIndex(path, "/")
	path = path[0 : i+1]
	c.GET(path).Handle(func(c *gin.Context) { inst.handle(c) })
}

func (inst *myStaticResHolder) handle(c *gin.Context) {
	data, err := inst.getContent()
	if err != nil {
		text := err.Error()
		c.JSON(500, text)
		return
	}
	charset := ";charset=utf-8"
	c.Data(200, inst.contentType+charset, data)
	// size := len(data)
	// fmt.Println("send-content-length: ", size)
}

func (inst *myStaticResHolder) wantCache(data []byte) bool {
	size := len(data)
	return (size < 1024*1024)
}

func (inst *myStaticResHolder) getContent() ([]byte, error) {
	data := inst.content
	if data == nil {
		data2, err := inst.loadContent()
		if err != nil {
			return nil, err
		}
		if inst.wantCache(data2) {
			inst.content = data2
		}
		data = data2
	}
	return data, nil
}

func (inst *myStaticResHolder) loadContent() ([]byte, error) {
	return inst.parent.loadContent(inst)
}

////////////////////////////////////////////////////////////////////////////////
