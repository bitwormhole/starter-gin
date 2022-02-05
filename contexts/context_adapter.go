package contexts

import (
	"context"

	"github.com/bitwormhole/starter/contexts"
	"github.com/gin-gonic/gin"
)

////////////////////////////////////////////////////////////////////////////////

type ginContextAdapter struct {
	gc *gin.Context
}

func (inst *ginContextAdapter) _Impl() contexts.ContextSetter {
	return inst
}

func (inst *ginContextAdapter) bind() {
	const key = keyGinContextBinding
	contexts.SetupContextSetter(inst)
	inst.SetValue(key, inst)
}

func (inst *ginContextAdapter) unbind() {
	inst.gc = nil
}

func (inst *ginContextAdapter) isReady() bool {
	return inst.gc != nil
}

func (inst *ginContextAdapter) GetContext() context.Context {
	return inst.gc
}

func (inst *ginContextAdapter) SetValue(key interface{}, value interface{}) {
	name := stringifyKeyName(key)
	inst.gc.Set(name, value)
}
