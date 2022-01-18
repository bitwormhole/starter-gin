package glass

import (
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// StaticWebResourcesController ...
type StaticWebResourcesController struct {
	markup.Component `class:"static-web-controller"`

	// 这 Root 是一个资源路径，指向 static-www 文件夹       `inject:"${web.static.root}"`
	Root      string              `inject:"${web.static.root}"`
	Container Container           `inject:"#gin-web-container"`
	Context   application.Context `inject:"context"`
}

func (inst *StaticWebResourcesController) _Impl() Controller {
	return inst
}

// Init ...
func (inst *StaticWebResourcesController) Init(ec EngineConnection) error {

	root := inst.Root
	app := inst.Context
	res := app.GetResources()
	list := res.List(root, true)

	for _, item := range list {

		if item.IsDir {
			continue
		}

		holder := &staticWebResourcesHolder{}
		holder.init()
		holder.resPath = item.AbsolutePath
		holder.resSource = res
		holder.contentType = inst.findContentType(item.RelativePath)

		handler := holder.getHandler()

		path := "./" + item.RelativePath
		ec.Handle("GET", path, handler)
		if inst.isIndex(path) {
			indexPath := inst.toIndexParentPath(path)
			ec.Handle("GET", indexPath, handler)
		}
	}

	return nil
}

func (inst *StaticWebResourcesController) isIndex(path string) bool {
	list := inst.Container.GetIndexPageNames()
	for _, name := range list {
		if strings.HasSuffix(path, "/"+name) {
			return true
		}
	}
	return false
}

func (inst *StaticWebResourcesController) toIndexParentPath(path string) string {
	index := strings.LastIndex(path, "/")
	if index < 0 {
		return "/"
	}
	return path[0 : index+1]
}

func (inst *StaticWebResourcesController) findContentType(path string) string {
	types := inst.Container.GetContentTypes()
	return types.Find(path)
}

////////////////////////////////////////////////////////////////////////////////

// staticWebResourcesHolder 是单个资源的 holder
type staticWebResourcesHolder struct {
	resSource         collection.Resources
	resPath           string
	maxLengthForCache int

	contentType string
	length      int
	name        string
	data        []byte

	fn gin.HandlerFunc
}

func (inst *staticWebResourcesHolder) init() {
	inst.maxLengthForCache = 1024 * 64
	inst.fn = func(c *gin.Context) { inst.handle(c) }
}

func (inst *staticWebResourcesHolder) getHandler() gin.HandlerFunc {
	return inst.fn
}

func (inst *staticWebResourcesHolder) loadData() ([]byte, error) {
	return inst.resSource.GetBinary(inst.resPath)
}

func (inst *staticWebResourcesHolder) getData() ([]byte, error) {
	data := inst.data
	if data == nil {
		data2, err := inst.loadData()
		if err != nil {
			return nil, err
		}
		data = data2
		size := len(data)
		inst.length = size
		if size <= inst.maxLengthForCache {
			inst.data = data
		}
	}
	return data, nil
}

func (inst *staticWebResourcesHolder) handle(c *gin.Context) {
	code := 200
	data, err := inst.getData()
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	c.Data(code, inst.contentType+"; charset=utf-8", data)
}

////////////////////////////////////////////////////////////////////////////////
