package web

import "github.com/gin-gonic/gin"

type GinContextHandlerFunc func(*gin.Context)

type Container interface {
	Engine() *gin.Engine

	Mapping(path string) Container

	AddFilter() FilterRegistration

	GET(path string) HandlerRegistration
	POST(path string) HandlerRegistration
	PUT(path string) HandlerRegistration
	DELETE(path string) HandlerRegistration
	Handle(method string, path string) HandlerRegistration
}
