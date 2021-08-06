package glass

import "github.com/gin-gonic/gin"

type ErrorController struct {

	// public
	ResourcePath string
	ContentType  string
	Status       int

	Container *Container

	// private
	data []byte
}

func (inst *ErrorController) _Impl() Controller {
	return inst
}

func (inst *ErrorController) Init(ec EngineConnection) error {

	err := inst.load()
	if err != nil {
		return err
	}

	ec.HandleNoMethod(func(c *gin.Context) { inst.handle(c) })
	ec.HandleNoResource(func(c *gin.Context) { inst.handle(c) })
	return nil
}

func (inst *ErrorController) load() error {
	path := inst.ResourcePath
	res := inst.Container.AppContext.GetResources()
	data, err := res.GetBinary(path)
	if err != nil {
		return err
	}
	inst.data = data
	return nil
}

func (inst *ErrorController) handle(c *gin.Context) {
	c.Data(inst.Status, inst.ContentType+"; charset=utf-8", inst.data)
}
