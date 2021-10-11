package glass

import (
	"context"
	"time"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
	"github.com/gin-gonic/gin"
)

const ginContextHolderKey = "glass.ginContextAdapter.1d18341c61c0ef62f7ac7bd.v1"

// CommonContext 函数把【*gin.Context】转换成【context.Context】
func CommonContext(c *gin.Context) context.Context {
	return EditableContext(c)
}

// EditableContext 函数把【*gin.Context】转换成【lang.Context】
func EditableContext(c *gin.Context) lang.Context {

	const key = ginContextHolderKey

	// try get older
	obj, ok1 := c.Get(key)
	if ok1 {
		holder, ok2 := obj.(*ginContextHolder)
		if ok2 && holder != nil {
			return holder.lc
		}
	}

	// make new
	holder := &ginContextHolder{}
	holder.lc = holder
	holder.gc = c
	c.Set(key, holder)
	holder.bind()
	return holder.lc
}

// GinContext 函数把【context.Context】转换成【*gin.Context】
func GinContext(cc context.Context) *gin.Context {

	const key = ginContextHolderKey
	obj := cc.Value(key)

	if obj != nil {
		holder, ok2 := obj.(*ginContextHolder)
		if ok2 && holder != nil {
			return holder.gc
		}
	}

	panic("The context.Context is not a binding to gin.Context")
}

////////////////////////////////////////////////////////////////////////////////

type ginContextHolder struct {
	// lang.Context

	lc lang.Context
	gc *gin.Context
	ac application.Context

	err error
}

func (inst *ginContextHolder) _Impl() lang.Context {
	return inst
}

func (inst *ginContextHolder) bind() {
	lc := inst.lc
	if lc != nil {
		lang.SetupContext(lc)
	}
}

func (inst *ginContextHolder) unbind() {
	inst.ac = nil
	inst.gc = nil
	inst.lc = nil
}

func (inst *ginContextHolder) GetValue(key string) interface{} {
	obj, ok := inst.gc.Get(key)
	if ok {
		return obj
	}
	return nil
}

func (inst *ginContextHolder) SetValue(key string, value interface{}) {
	inst.gc.Set(key, value)
}

func (inst *ginContextHolder) Deadline() (deadline time.Time, ok bool) {
	t := time.Now()
	return t, false
}

func (inst *ginContextHolder) Done() <-chan struct{} {
	return nil
}

func (inst *ginContextHolder) Err() error {
	return inst.err
}

func (inst *ginContextHolder) Value(key interface{}) interface{} {
	if key == nil {
		return nil
	}
	strkey := key.(string)
	v, ok := inst.gc.Get(strkey)
	if ok {
		return v
	}
	return nil
}
