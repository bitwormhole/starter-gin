package glass

import "github.com/gin-gonic/gin"

// Interceptor ...
type Interceptor interface {

	// input: 目标handler , output: 目标的代理
	Intercept(h gin.HandlerFunc) gin.HandlerFunc
}

// InterceptorRegistration ...
type InterceptorRegistration struct {
	Name        string
	Order       int
	Interceptor Interceptor
}

// InterceptorRegistry ...
// 通过【inject:".rest-interceptor-registry"】
// 或者【inject:".static-web-interceptor-registry"】注入
type InterceptorRegistry interface {
	GetRegistrationList() []*InterceptorRegistration
}
