package glass

import "github.com/gin-gonic/gin"

type StaticWebResourcesController struct {
}

func (inst *StaticWebResourcesController) _Impl() Controller {
	return inst
}

func (inst *StaticWebResourcesController) Init(ec EngineConnection) error {

	ec.Handle("GET", "mock-static-res-1", func(c *gin.Context) { inst.h(c) })

	return nil
}

func (inst *StaticWebResourcesController) h(c *gin.Context) {
	c.JSON(200, gin.H{
		"a": "1",
		"b": "2",
		"c": "3",
	})
}
