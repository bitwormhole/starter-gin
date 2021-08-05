package demo

import (
	"time"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type theExampleController1 struct {
	markup.Component
	instance *ExampleController1 `class:"rest-controller"`
}

////////////////////////////////////////////////////////////////////////////////

type ExampleController1 struct {
}

func (inst *ExampleController1) _Impl() glass.Controller {
	return inst
}

func (inst *ExampleController1) Init(conn glass.EngineConnection) error {

	conn = conn.RequestMapping("example1")

	conn.Handle("GET", ":id", func(c *gin.Context) { inst.doGet(c) })
	conn.Handle("GET", "", func(c *gin.Context) { inst.doGet(c) })
	conn.Handle("POST", "", func(c *gin.Context) { inst.doPost(c) })
	conn.Handle("PUT", ":id", func(c *gin.Context) { inst.doPut(c) })
	conn.Handle("DELETE", ":id", func(c *gin.Context) { inst.doDelete(c) })

	return nil
}

func (inst *ExampleController1) doDelete(c *gin.Context) {
	inst.h(c)
}

func (inst *ExampleController1) doGet(c *gin.Context) {
	inst.h(c)
}

func (inst *ExampleController1) doPost(c *gin.Context) {
	inst.h(c)
}

func (inst *ExampleController1) doPut(c *gin.Context) {
	inst.h(c)
}

func (inst *ExampleController1) h(c *gin.Context) {

	now := time.Now()
	t := now.String()
	method := c.Request.Method
	id := c.Param("id")

	c.JSON(200, gin.H{
		"Method": method,
		"ID":     id,
		"Time":   t,
	})
}
