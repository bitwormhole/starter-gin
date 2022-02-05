package contexts

import (
	"context"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

const keyGinContextBinding = "glass.ginContextAdapter.1d18341c61c0ef62f7ac7bd.v1"

// GetContext2 函数把【*gin.Context】转换成【context.Context】
func GetContext2(c *gin.Context) (context.Context, error) {
	return c, nil
}

// GetGinContext 函数把【context.Context】转换成【*gin.Context】
func GetGinContext(cc context.Context) (*gin.Context, error) {
	const key = keyGinContextBinding
	o1 := cc.Value(key)
	o2, ok := o1.(*ginContextAdapter)
	if ok {
		gc := o2.gc
		if gc != nil {
			return gc, nil
		}
	}
	return nil, errors.New("The context.Context is not a binding to gin.Context")
}

// SetupGinContext 函数根据 gin.Context 初始化一个绑定的 context.Context
func SetupGinContext(c *gin.Context) error {
	const key = keyGinContextBinding
	o1 := c.Value(key)
	o2, ok := o1.(*ginContextAdapter)
	if ok {
		if o2.isReady() {
			return nil
		}
	}
	o2 = &ginContextAdapter{gc: c}
	o2.bind()
	return nil
}

func stringifyKeyName(key interface{}) string {
	name, ok := key.(string)
	if ok {
		return name
	}
	return fmt.Sprint(key)
}
