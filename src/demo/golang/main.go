package main

import (
	"time"

	"github.com/bitwormhole/starter"
	ginstarter "github.com/bitwormhole/starter-gin"
	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/vlog"
)

func main() {
	i := starter.InitApp()
	i.Use(ginstarter.ModuleForDemo())
	rt, err := i.RunEx()
	if err != nil {
		panic(err)
	}

	stopped := false
	go func() {
		rt.Loop()
		stopped = true
	}()

	const selector = "#gin-web-container"
	com, err := rt.Context().GetComponent(selector)
	if err != nil {
		panic(err)
	}

	for tout := 5; tout > 0; tout-- {
		vlog.Info("shutdown in ", tout, "sec ...")
		time.Sleep(time.Second)
	}

	container, ok := com.(glass.Container)
	if ok && container != nil {
		container.Shutdown()
	}

	for !stopped {
		time.Sleep(time.Second)
	}
}
