package web

import "github.com/gin-gonic/gin"

type Http404pageController struct {
}

func (inst *Http404pageController) __impl__() Controller {
	return inst
}

func (inst *Http404pageController) Config(c Container) error {

	c.HandleNoMethod(func(c *gin.Context) { inst.http404(c) })
	c.HandleNoResource(func(c *gin.Context) { inst.http404(c) })
	return nil
}

func (inst *Http404pageController) http404(c *gin.Context) {

	p := map[string]string{
		"method":  c.Request.Method,
		"path":    c.Request.URL.Path,
		"message": "NOT FOUND",
		"status":  "404",
	}

	c.JSON(404, p)
}
