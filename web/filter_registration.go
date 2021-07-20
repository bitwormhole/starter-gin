package web

import "github.com/gin-gonic/gin"

type FilterRegistration interface {

	// priority 数字越大，优先级越高，越先处理
	Priority(priority int) FilterRegistration

	Container() Container

	Handle(h gin.HandlerFunc)
}
