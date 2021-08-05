package web

import (
	"strings"

	"github.com/bitwormhole/starter/collection"
)

// ContentTypeManager 是web内容类型管理器
type ContentTypeManager struct {
	inner         *innerContentTypeManager
	ServerContext *GinServerContext
}

func (inst *ContentTypeManager) getInner() *innerContentTypeManager {
	inner := inst.inner
	if inner == nil {
		inner := &innerContentTypeManager{}
		inner.init(inst.ServerContext)
		inst.inner = inner
	}
	return inner
}

// FindContentTypeBySuffix 根据文件名后缀查找对应的类型
func (inst *ContentTypeManager) FindContentTypeBySuffix(suffix string) string {
	index := strings.LastIndex(suffix, ".")
	if index >= 0 {
		suffix = suffix[index+1:]
	} else {
		// NOP
	}
	key := "types." + strings.ToLower(suffix)
	return inst.getInner().props.GetProperty(key, "application/octet-stream")
}

////////////////////////////////////////////////////////////////////////////////

type innerContentTypeManager struct {
	props collection.Properties
}

func (inst *innerContentTypeManager) init(sc *GinServerContext) {
	cfg := sc.Configuration
	app := sc.ApplicationContext
	res := app.GetResources()
	uri := cfg.ContentTypeProperties
	props, err := inst.loadPropertiesFromRes(res, uri)
	if err != nil {
		props = inst.loadPropertiesDefault()
	}
	inst.props = props
}

func (inst *innerContentTypeManager) loadPropertiesFromRes(res collection.Resources, uri string) (collection.Properties, error) {
	text, err := res.GetText(uri)
	if err != nil {
		return nil, err
	}
	return collection.ParseProperties(text, nil)
}

func (inst *innerContentTypeManager) loadPropertiesDefault() collection.Properties {
	const prefix = "types."
	props := collection.CreateProperties()

	props.SetProperty(prefix+"md", "text/markdown")
	props.SetProperty(prefix+"txt", "text/plain")
	props.SetProperty(prefix+"htm", "text/html")
	props.SetProperty(prefix+"html", "text/html")
	props.SetProperty(prefix+"css", "text/css")

	props.SetProperty(prefix+"js", "application/x-javascript")

	props.SetProperty(prefix+"jpg", "image/jpeg")
	props.SetProperty(prefix+"gif", "image/gif")
	props.SetProperty(prefix+"png", "image/png")
	props.SetProperty(prefix+"ico", "image/x-icon")
	props.SetProperty(prefix+"tif", "image/tiff")
	props.SetProperty(prefix+"tiff", "image/tiff")
	props.SetProperty(prefix+"bmp", "application/x-bmp")

	return props
}
