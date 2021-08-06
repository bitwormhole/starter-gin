package demo

import (
	"time"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type theExampleController3 struct {
	markup.Component
	instance *ExampleController3 `class:"rest-controller"`
}

////////////////////////////////////////////////////////////////////////////////

type ExampleController3 struct {
}

func (inst *ExampleController3) _Impl() glass.Controller {
	return inst
}

func (inst *ExampleController3) Init(conn glass.EngineConnection) error {

	conn = conn.RequestMapping("example3")

	conn.Handle("GET", "", func(c *gin.Context) { inst.h(c) })
	conn.Handle("GET", "./:id/x/y/z", func(c *gin.Context) { inst.h(c) })
	conn.Handle("POST", "", func(c *gin.Context) { inst.h(c) })

	// conn.Filter(1, func(c *gin.Context) { inst.h(c) })
	// conn.Filter(2, func(c *gin.Context) { inst.h(c) })
	// conn.Filter(3, func(c *gin.Context) { inst.h(c) })

	return nil
}

func (inst *ExampleController3) h(c *gin.Context) {

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
