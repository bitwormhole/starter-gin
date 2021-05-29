package filters

import (
	"github.com/bitwormhole/starter-gin/security"
	"github.com/bitwormhole/starter-gin/web/containers"
	"github.com/gin-gonic/gin"
)

type SecuritySessionFilter struct {
	sessionFactory security.SessionFactory
}

func (inst *SecuritySessionFilter) _impl_filter() containers.GinWebFilter {
	return inst
}

func (inst *SecuritySessionFilter) SetSessionFactory(f security.SessionFactory) {
	inst.sessionFactory = f
}

func (inst *SecuritySessionFilter) _do_filter_(c *gin.Context) error {
	factory := inst.sessionFactory
	session, err := security.OpenSession(c, factory)
	if err != nil {
		// return err
	}
	c.Next()
	session.Close()
	return nil
}

func (inst *SecuritySessionFilter) Priority() int {
	return 0
}

func (inst *SecuritySessionFilter) ConfigFilter(engine *gin.Engine) error {
	engine.Use(func(c *gin.Context) {
		inst._do_filter_(c)
	})
	return nil
}
