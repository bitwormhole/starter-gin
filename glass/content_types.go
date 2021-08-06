package glass

import (
	"log"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

type ContentTypeManager interface {
	Find(pathname string) string
}

////////////////////////////////////////////////////////////////////////////////

type DefaultContentTypeManager struct {
	AppContext      application.Context
	TypesProperties string

	// private
	table collection.Properties
}

func (inst *DefaultContentTypeManager) _Impl() ContentTypeManager {
	return inst
}

func (inst *DefaultContentTypeManager) loadTable() (collection.Properties, error) {
	res := inst.AppContext.GetResources()
	text, err := res.GetText(inst.TypesProperties)
	if err != nil {
		return nil, err
	}
	return collection.ParseProperties(text, nil)
}

func (inst *DefaultContentTypeManager) loadDefaultTable() collection.Properties {

	log.Println("[INFO] load default content-types in code")

	const prefix = "content-types."
	table := collection.CreateProperties()

	table.SetProperty(prefix+"txt", "text/plain")
	table.SetProperty(prefix+"htm", "text/html")
	table.SetProperty(prefix+"html", "text/html")

	table.SetProperty(prefix+"jpg", "image/jpg")
	table.SetProperty(prefix+"gif", "image/gif")
	table.SetProperty(prefix+"png", "image/png")
	table.SetProperty(prefix+"ico", "image/x-icon")

	table.SetProperty(prefix+"js", "application/x-javascript")
	table.SetProperty(prefix+"css", "text/css")

	return table
}

func (inst *DefaultContentTypeManager) getTable() collection.Properties {
	table := inst.table
	if table == nil {
		t2, err := inst.loadTable()
		if err != nil {
			log.Println("[WARNING] cannot load content-types:" + err.Error())
			table = inst.loadDefaultTable()
		} else {
			table = t2
		}
		inst.table = table
	}
	return table
}

func (inst *DefaultContentTypeManager) keyFor(name string) string {
	const prefix = "content-types"
	index := strings.LastIndex(name, ".")
	if index < 0 {
		return ""
	}
	key := prefix + name[index:]
	return strings.ToLower(key)
}

// Find 根据文件名查找相应的content-type
func (inst *DefaultContentTypeManager) Find(name string) string {
	const defaultType = "application/octet-stream"
	key := inst.keyFor(name)
	return inst.getTable().GetProperty(key, defaultType)
}

////////////////////////////////////////////////////////////////////////////////
