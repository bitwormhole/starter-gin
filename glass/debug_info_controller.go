package glass

import (
	"net/http"

	"github.com/bitwormhole/starter/application"
	"github.com/gin-gonic/gin"
)

// DebugInfoController 是一个REST控制器，用来显示调试信息
type DebugInfoController struct {
	AppContext application.Context
}

func (inst *DebugInfoController) _Impl() Controller {
	return inst
}

func (inst *DebugInfoController) Init(e EngineConnection) error {
	e = e.RequestMapping("debug")
	e.Handle(http.MethodGet, "", func(c *gin.Context) { inst.handleGet(c) })
	return nil
}

func (inst *DebugInfoController) handleGet(c *gin.Context) {

	app := inst.AppContext
	props := app.GetProperties().Export(nil)
	env := app.GetEnvironment().Export(nil)
	args := app.GetArguments().Export()

	h := gin.H{}
	h["arguments"] = args
	h["environment"] = env
	h["properties"] = props
	c.JSON(200, h)
}
