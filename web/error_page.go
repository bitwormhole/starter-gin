package web

import (
	"github.com/gin-gonic/gin"
)

type Http404pageController struct {
	ServerContext *GinServerContext

	// private
	noMethodPage   StaticResourceHolder
	noMethodStatus int
	noResPage      StaticResourceHolder
	noResStatus    int
}

func (inst *Http404pageController) __impl__() Controller {
	return inst
}

func (inst *Http404pageController) init() error {

	cfg := inst.ServerContext.Configuration
	noMethodPage, err := inst.loadStaticPage(cfg.NoMethodPage)
	if err != nil {
		return err
	}

	noResPage, err := inst.loadStaticPage(cfg.NoResourcePage)
	if err != nil {
		return err
	}

	inst.noMethodPage = noMethodPage
	inst.noMethodStatus = cfg.NoMethodStatus
	inst.noResPage = noResPage
	inst.noResStatus = cfg.NoResourceStatus
	return nil
}

func (inst *Http404pageController) loadStaticPage(name string) (StaticResourceHolder, error) {
	loader := inst.ServerContext.ResourceLoader
	holder := loader.Load(name)
	err := holder.Error()
	return holder, err
}

func (inst *Http404pageController) Config(c Container) error {

	err := inst.init()
	if err != nil {
		return err
	}

	c.HandleNoMethod(func(c *gin.Context) { inst.handleNoMethod(c) })
	c.HandleNoResource(func(c *gin.Context) { inst.handleNoResource(c) })
	return nil
}

func (inst *Http404pageController) handleNoMethod(c *gin.Context) {

	code := inst.noMethodStatus
	page := inst.noMethodPage

	contentType := page.ContentType()
	data := page.Data()
	c.Data(code, contentType, data)
}

func (inst *Http404pageController) handleNoResource(c *gin.Context) {

	code := inst.noResStatus
	page := inst.noResPage

	contentType := page.ContentType()
	data := page.Data()
	c.Data(code, contentType, data)
}
