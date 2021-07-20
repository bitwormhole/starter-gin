package web

import "github.com/gin-gonic/gin"

type HandlerRegistration interface {
	Method(method string) HandlerRegistration

	Mapping(path string) HandlerRegistration

	Container() Container

	Handle(h gin.HandlerFunc)
}
