package contexts

import (
	"context"
	"time"

	"github.com/bitwormhole/starter/contexts"
	"github.com/gin-gonic/gin"
)

////////////////////////////////////////////////////////////////////////////////

type ginContextAdapter struct {
	gc       *gin.Context
	err      error
	deadline time.Time
}

func (inst *ginContextAdapter) _Impl() (context.Context, contexts.ContextSetter) {
	return inst, inst
}

func (inst *ginContextAdapter) bind() {
	const key = keyGinContextBinding
	contexts.SetupContextSetter(inst)
	inst.SetValue(key, inst)
}

func (inst *ginContextAdapter) unbind() {
	inst.err = nil
	inst.gc = nil
}

func (inst *ginContextAdapter) isReady() bool {
	return inst.gc != nil
}

// func (inst *ginContextAdapter) GetValue(key string) interface{} {
// 	obj, ok := inst.gc.Get(key)
// 	if ok {
// 		return obj
// 	}
// 	return nil
// }

// func (inst *ginContextAdapter) SetValue(key string, value interface{}) {
// 	inst.gc.Set(key, value)
// }

func (inst *ginContextAdapter) Deadline() (deadline time.Time, ok bool) {
	return inst.deadline, false
}

func (inst *ginContextAdapter) Done() <-chan struct{} {
	return nil
}

func (inst *ginContextAdapter) Err() error {
	return inst.err
}

func (inst *ginContextAdapter) Value(key interface{}) interface{} {
	name := stringifyKeyName(key)
	v, ok := inst.gc.Get(name)
	if ok {
		return v
	}
	return nil
}

func (inst *ginContextAdapter) GetContext() context.Context {
	return inst
}

func (inst *ginContextAdapter) SetValue(key interface{}, value interface{}) {
	name := stringifyKeyName(key)
	inst.gc.Set(name, value)
}
