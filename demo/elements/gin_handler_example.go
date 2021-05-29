package elements

import (
	"github.com/bitwormhole/starter-gin/security"
	"github.com/bitwormhole/starter-gin/web/containers"
	"github.com/bitwormhole/starter-gin/web/rest"
	"github.com/gin-gonic/gin"
)

type ExampleGinController struct {
	REST rest.Context
}

func (inst *ExampleGinController) _impl_controller() containers.GinWebController {
	return inst
}

func (inst *ExampleGinController) Config(engine *gin.Engine) error {

	rt := inst.REST.ForType("example1")

	engine.GET(rt.Path(""), func(c *gin.Context) { inst.doGet(c) })
	engine.POST(rt.Path(""), func(c *gin.Context) { inst.doPost(c) })

	engine.GET(rt.Path(":id"), func(c *gin.Context) { inst.doGet(c) })
	engine.PUT(rt.Path(":id"), func(c *gin.Context) { inst.doPut(c) })
	engine.DELETE(rt.Path(":id"), func(c *gin.Context) { inst.doDelete(c) })

	return nil
}

func (inst *ExampleGinController) doGet(c *gin.Context) {

	id := c.Param("id")

	table := map[string]string{
		"hello":  "world",
		"method": "get",
		"type":   "example1",
		"id":     id,
	}

	session, err := security.GetSession(c)
	if err != nil {
		c.Error(err)
		return
	}
	session.Commit()
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
