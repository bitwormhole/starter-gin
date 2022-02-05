package factory

import (
	"github.com/bitwormhole/starter-gin/contexts"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// ContextBindController 控制器把 context.Context 绑定到 gin.Context
type ContextBindController struct {
	markup.Component `class:"rest-controller"`

	// AppContext application.Context `inject:"context"`

	Order int `inject:"${webfilter.context.order}"`
}

func (inst *ContextBindController) _Impl() glass.Controller {
	return inst
}

// Init 初始化控制器
func (inst *ContextBindController) Init(ec glass.EngineConnection) error {
	ec.Filter(inst.Order, func(c *gin.Context) { inst.doFilter(c) })
	return nil
}

func (inst *ContextBindController) doFilter(c *gin.Context) {
	contexts.SetupGinContext(c)
	c.Next()
}

////////////////////////////////////////////////////////////////////////////////
