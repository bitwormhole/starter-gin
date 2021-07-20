package elements

import (
	"github.com/bitwormhole/starter-gin/web"
	"github.com/bitwormhole/starter/application"
	"github.com/gin-gonic/gin"
)

type ExampleGinController struct {
}

func (inst *ExampleGinController) _impl_controller() web.Controller {
	return inst
}

func (inst *ExampleGinController) Inject(ctx application.Context) error {
	return nil
}

func (inst *ExampleGinController) Init() error {
	return nil
}

func (inst *ExampleGinController) Destroy() error {
	return nil
}

func (inst *ExampleGinController) Config(con web.Container) error {

	con = con.Mapping("example1")

	con.AddFilter().Priority(8).Handle(func(c *gin.Context) { inst.doFilter(c) })

	con.GET("").Handle(func(c *gin.Context) { inst.doGet(c) })
	con.POST("").Handle(func(c *gin.Context) { inst.doGet(c) })

	con.GET(":id").Handle(func(c *gin.Context) { inst.doGet(c) })
	con.PUT(":id").Handle(func(c *gin.Context) { inst.doPut(c) })
	con.DELETE(":id").Handle(func(c *gin.Context) { inst.doDelete(c) })

	return nil
}

func (inst *ExampleGinController) doFilter(c *gin.Context) {
	c.Next()
}

func (inst *ExampleGinController) doGet(c *gin.Context) {

	id := c.Param("id")

	table := map[string]string{
		"hello":  "world",
		"method": "get",
		"type":   "example1",
		"id":     id,
	}

	c.JSON(200, table)
}

func (inst *ExampleGinController) doPost(c *gin.Context) {
	table := map[string]string{
		"hello":  "world",
		"method": "post",
	}
	c.JSON(200, table)
}

func (inst *ExampleGinController) doPut(c *gin.Context) {
	table := map[string]string{
		"hello":  "world",
		"method": "put",
	}
	c.JSON(200, table)
}

func (inst *ExampleGinController) doDelete(c *gin.Context) {
	table := map[string]string{
		"hello":  "world",
		"method": "delete",
	}
	c.JSON(200, table)
}
