package devtools

import (
	"reflect"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
	"github.com/gin-gonic/gin"
)

type apiComponentsHandler struct {
	all map[string]application.ComponentHolder
}

func (inst *apiComponentsHandler) init(ac application.Context) {
	inst.all = ac.GetComponents().Export(nil)
}

func (inst *apiComponentsHandler) handle(c *gin.Context) {
	list := make([]*ComponentInfoDTO, 0)
	for _, holder := range inst.all {
		dto := inst.makeDTO(holder)
		list = append(list, dto)
	}
	h := gin.H{
		"components": list,
	}
	c.JSON(200, h)
}

func (inst *apiComponentsHandler) makeDTO(h application.ComponentHolder) *ComponentInfoDTO {

	info := h.GetInfo()

	dto := &ComponentInfoDTO{}
	dto.Class = info.GetClasses()
	dto.ID = info.GetID()
	dto.Scope = inst.stringifyScope(info.GetScope())
	dto.Type = inst.stringifyPrototype(info.GetPrototype())

	return dto
}

func (inst *apiComponentsHandler) stringifyPrototype(o lang.Object) string {
	return reflect.ValueOf(o).String()
}

func (inst *apiComponentsHandler) stringifyScope(scope application.ComponentScope) string {
	switch scope {
	case application.ScopeContext:
		return "context"
	case application.ScopeSingleton:
		return "singleton"
	case application.ScopePrototype:
		return "prototype"
	default:
		return "unknow"
	}
}
