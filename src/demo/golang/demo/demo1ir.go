package demo

import (
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
	"github.com/gin-gonic/gin"
)

type Demo1ir struct {
	markup.Component `class:"rest-interceptor-registry"`
}

func (inst *Demo1ir) _Impl() glass.InterceptorRegistry {
	return inst
}

func (inst *Demo1ir) GetRegistrationList() []*glass.InterceptorRegistration {

	ir1 := &glass.InterceptorRegistration{
		Name:        "ir1",
		Order:       1,
		Interceptor: &Interceptor1{},
	}

	ir2 := &glass.InterceptorRegistration{
		Name:        "ir2",
		Order:       2,
		Interceptor: &Interceptor2{},
	}

	return []*glass.InterceptorRegistration{ir1, ir2}
}

////////////////////////////////////////////////////////////////////////////////

type Interceptor1 struct {
}

func (inst *Interceptor1) _Impl() glass.Interceptor {
	return inst
}

func (inst *Interceptor1) Intercept(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c)
	}
}

////////////////////////////////////////////////////////////////////////////////

type Interceptor2 struct {
}

func (inst *Interceptor2) _Impl() glass.Interceptor {
	return inst
}

func (inst *Interceptor2) Intercept(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h(c)
	}
}

////////////////////////////////////////////////////////////////////////////////
