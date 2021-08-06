package demo

import (
	"log"
	"time"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type theExampleController2 struct {
	markup.Component
	instance *ExampleController2 `class:"rest-controller"`
}

////////////////////////////////////////////////////////////////////////////////

type ExampleController2 struct {
}

func (inst *ExampleController2) _Impl() glass.Controller {
	return inst
}

func (inst *ExampleController2) Init(conn glass.EngineConnection) error {

	conn = conn.RequestMapping("example2")

	conn.Handle("GET", "", func(c *gin.Context) { inst.h(c) })
	conn.Handle("POST", "", func(c *gin.Context) { inst.h(c) })

	conn.Filter(3, func(c *gin.Context) { inst.f3(c) })
	conn.Filter(1, func(c *gin.Context) { inst.f1(c) })
	conn.Filter(2, func(c *gin.Context) { inst.f2(c) })

	return nil
}

func (inst *ExampleController2) f1(c *gin.Context) {
	log.Println("do:filter:1")
	c.Next()
}

func (inst *ExampleController2) f2(c *gin.Context) {
	log.Println("do:filter:2")
	c.Next()
}

func (inst *ExampleController2) f3(c *gin.Context) {
	log.Println("do:filter:3")
	c.Next()
}

func (inst *ExampleController2) h(c *gin.Context) {

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
