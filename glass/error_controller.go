package glass

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

// ErrorController ...
type ErrorController struct {
	markup.Component `class:"static-web-controller"`

	// public

	ResourcePath string              `inject:"${web.error-page.resource}"`
	ContentType  string              `inject:"${web.error-page.content-type}"`
	Status       int                 `inject:"${web.error-page.status}"`
	Container    Container           `inject:"#gin-web-container"`
	Context      application.Context `inject:"context"`

	// private
	data []byte
}

func (inst *ErrorController) _Impl() Controller {
	return inst
}

// Init ...
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
	res := inst.Context.GetResources()
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
