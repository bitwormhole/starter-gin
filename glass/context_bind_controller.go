package glass

import (
	"errors"

	"github.com/bitwormhole/starter/application"
	"github.com/gin-gonic/gin"
)

// ContextBindController 控制器把 application.Context 绑定到 gin.Context
type ContextBindController struct {
	AppContext application.Context
}

func (inst *ContextBindController) _Impl() Controller {
	return inst
}

// Init 初始化控制器
func (inst *ContextBindController) Init(ec EngineConnection) error {
	ec.Filter(0, func(c *gin.Context) { inst.doFilter(c) })
	return nil
}

func (inst *ContextBindController) doFilter(c *gin.Context) {
	holder := getWebApplicationContextHolder(c)
	holder.ac = inst.AppContext
	defer func() { holder.ac = nil }()
	c.Next()
}

////////////////////////////////////////////////////////////////////////////////

type webApplicationContextHolder struct {
	ac application.Context
}

const keyForWebApplicationContextHolder = "webApplicationContextHolder.6644c5bff5b55acc932c0f576f#binding"

func getWebApplicationContextHolder(c *gin.Context) *webApplicationContextHolder {
	const key = keyForWebApplicationContextHolder
	var holder *webApplicationContextHolder = nil
	obj, ok := c.Get(key)
	if ok {
		h2, ok := obj.(*webApplicationContextHolder)
		if ok {
			holder = h2
		}
	}
	if holder == nil {
		holder = &webApplicationContextHolder{}
		c.Set(key, holder)
	}
	return holder
}

////////////////////////////////////////////////////////////////////////////////

// WebApplicationContext 函数把【*gin.Context】转换成【application.Context】
func WebApplicationContext(c *gin.Context) (application.Context, error) {
	holder := getWebApplicationContextHolder(c)
	ac := holder.ac
	if ac == nil {
		return nil, errors.New("no application.Context has bound to the gin.Context")
	}
	return ac, nil
}
