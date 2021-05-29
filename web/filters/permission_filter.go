package filters

import (
	"github.com/bitwormhole/starter-gin/web/containers"
	"github.com/gin-gonic/gin"
)

type PermissionFilter struct {
	// sessionFactory security.SessionFactory
}

func (inst *PermissionFilter) _impl_filter() containers.GinWebFilter {
	return inst
}

func (inst *PermissionFilter) ConfigFilter(engine *gin.Engine) error {
	engine.Use(func(c *gin.Context) {
		inst.handle(c)
	})
	return nil
}

func (inst *PermissionFilter) Priority() int {
	return 0
}

func (inst *PermissionFilter) handle(c *gin.Context) error {

	// c.Next()

	return nil
}
