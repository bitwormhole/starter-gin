package demo

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter-restful/api/vo"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type Demo1ctrl struct {
	markup.Component `class:"rest-controller"`

	Resp glass.MainResponder `inject:"#glass-main-responder"`
}

func (inst *Demo1ctrl) _Impl() glass.Controller {
	return inst
}

func (inst *Demo1ctrl) Init(ec glass.EngineConnection) error {
	ec.Handle(http.MethodGet, "demo1", inst.handle)
	return nil
}

func (inst *Demo1ctrl) handle(c *gin.Context) {

	session := &vo.Session{}

	resp := glass.Response{}
	resp.Context = c
	resp.Data = session
	resp.Error = nil
	resp.Status = http.StatusOK

	inst.Resp.Send(&resp)
}
