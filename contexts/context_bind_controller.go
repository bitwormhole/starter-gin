package contexts

import (
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/application"
	"github.com/gin-gonic/gin"
)

// ContextBindController 控制器把 context.Context 绑定到 gin.Context
type ContextBindController struct {
	AppContext application.Context
}

func (inst *ContextBindController) _Impl() glass.Controller {
	return inst
}

// Init 初始化控制器
func (inst *ContextBindController) Init(ec glass.EngineConnection) error {
	ec.Filter(0, func(c *gin.Context) { inst.doFilter(c) })
	return nil
}

func (inst *ContextBindController) doFilter(c *gin.Context) {
	SetupGinContext(c)
	c.Next()
}

////////////////////////////////////////////////////////////////////////////////
