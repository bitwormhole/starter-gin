package devtools

import (
	"net/http"
	"sort"
	"time"

	"github.com/bitwormhole/starter-gin/contexts"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/application"
	contexts0 "github.com/bitwormhole/starter/contexts"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// DevtoolsController 是一个REST控制器，用来显示调试信息
type DevtoolsController struct {
	markup.Component `class:"rest-controller"`

	AppContext application.Context `inject:"context"`
	Enable     bool                `inject:"${web.devtools.enable}"`

	// private
	accumulator RequestAccumulator
}

func (inst *DevtoolsController) _Impl() glass.Controller {
	return inst
}

// Init 初始化控制器
func (inst *DevtoolsController) Init(e glass.EngineConnection) error {

	if !inst.Enable {
		return nil
	}

	err := inst.initFilter(e)
	if err != nil {
		return err
	}

	err = inst.initAPI(e)
	if err != nil {
		return err
	}

	err = inst.initStatic(e)
	if err != nil {
		return err
	}

	return nil
}

func (inst *DevtoolsController) initFilter(e glass.EngineConnection) error {
	e.Filter(0, func(c *gin.Context) { inst.doFilter(c) })
	return nil
}

func (inst *DevtoolsController) initAPI(e glass.EngineConnection) error {

	e = e.RequestMapping("/api/devtools.v1")

	e.Handle(http.MethodGet, "context", func(c *gin.Context) { inst.doGetContextInfo(c) })
	e.Handle(http.MethodGet, "requests", func(c *gin.Context) { inst.doGetRequestSum(c) })
	e.Handle(http.MethodDelete, "requests", func(c *gin.Context) { inst.doDeleteRequestSum(c) })
	e.Handle(http.MethodGet, "components", func(c *gin.Context) { inst.doGetComponents(c) })
	e.Handle(http.MethodGet, "modules", func(c *gin.Context) { inst.doGetModules(c) })

	e.Handle(http.MethodGet, "test/:id", func(c *gin.Context) { inst.doGetTest(c) })

	return nil
}

func (inst *DevtoolsController) initStatic(e glass.EngineConnection) error {
	e = e.RequestMapping("/debug")
	return nil
}

func (inst *DevtoolsController) now() int64 {
	return time.Now().Unix() * 1000
}

func (inst *DevtoolsController) doFilter(c *gin.Context) {

	method := c.Request.Method
	path := c.Request.URL.Path
	// log.Println("[debug]:   ", method, path)

	rec := &RequestRecordDTO{}
	rec.Method = method
	rec.Path = path
	rec.TimeBegin = inst.now()

	c.Next()

	rec.TimeEnd = inst.now()
	inst.accumulator.Add(rec)
}

func (inst *DevtoolsController) doGetContextInfo(c *gin.Context) {

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

func (inst *DevtoolsController) doGetComponents(c *gin.Context) {
	handler := &apiComponentsHandler{}
	handler.init(inst.AppContext)
	handler.handle(c)
}

func (inst *DevtoolsController) doGetModules(c *gin.Context) {
	handler := &apiModulesHandler{}
	handler.init(inst.AppContext)
	handler.handle(c)
}

func (inst *DevtoolsController) doGetTest(c *gin.Context) {

	cc, err := contexts.GetContext2(c)
	if err != nil {
		c.Error(err)
		return
	}

	setter, err := contexts0.GetContextSetter(cc)
	if err != nil {
		c.Error(err)
		return
	}

	setter.SetValue("foo", "bar")
	c.JSON(http.StatusOK, "ok")
}

func (inst *DevtoolsController) doDeleteRequestSum(c *gin.Context) {
	// reset
	inst.accumulator.reset()
}

func (inst *DevtoolsController) doGetRequestSum(c *gin.Context) {

	all := inst.accumulator.getEndpoints()
	keys := make([]string, 0)
	sumlist := make([]*RequestSumDTO, 0)
	h := gin.H{}

	for key := range all {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		item := all[key]
		sum := item.GetResult()
		sumlist = append(sumlist, sum)
	}

	h["requests"] = sumlist

	c.JSON(200, h)
}
