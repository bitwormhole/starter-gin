package demo

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type Demo1ctrl struct {
	markup.Component `class:"rest-controller"`
}

func (inst *Demo1ctrl) _Impl() glass.Controller {
	return inst
}

func (inst *Demo1ctrl) Init(ec glass.EngineConnection) error {
	ec.Handle(http.MethodGet, "demo1", inst.handle)
	return nil
}

func (inst *Demo1ctrl) handle(c *gin.Context) {

	c.JSON(http.StatusOK, "demo1: OK")
}
