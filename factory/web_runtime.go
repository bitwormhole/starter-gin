package factory

import "github.com/bitwormhole/starter-gin/glass"

type webRuntime struct {
	container glass.Container
}

func (inst *webRuntime) start() error {}

func (inst *webRuntime) stop() error {}

func (inst *webRuntime) loop() error {}
