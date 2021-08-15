package main

import (
	"net/http"

	ginstarter "github.com/bitwormhole/starter-gin"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/gin-gonic/gin"
)

func doGet(c *gin.Context) {

	path := c.FullPath()

	c.JSON(200, gin.H{
		"path": path,
	})
}

func services(ec glass.EngineConnection) error {

	ec = ec.RequestMapping("example1")
	ec.Handle(http.MethodGet, "", func(c *gin.Context) { doGet(c) })
	ec.Handle(http.MethodGet, ":id", func(c *gin.Context) { doGet(c) })

	return nil
}

func main() {

	mob := &ginstarter.GinModuleBuilder{}
	mob.Name("github.com/bitwormhole/starter-gin/+src/test/golang").Version("v1").Revision(1)
	mob.RegisterControllerFunc(services)

	mod := mob.Create()
	ginstarter.InitGin().Use(mod).Use(ginstarter.ModuleWithDevtools()).Run()
}
