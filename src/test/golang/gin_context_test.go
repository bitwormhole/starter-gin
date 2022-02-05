package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGinContext(t *testing.T) {
	engine := gin.Default()
	engine.Handle(http.MethodGet, "/test", func(c *gin.Context) {
		doTestGinContext1(c)
	})
}

func doTestGinContext1(c *gin.Context) {
	doTestGinContext2(c)
}

func doTestGinContext2(c context.Context) {
	c.Value("abc")
}
