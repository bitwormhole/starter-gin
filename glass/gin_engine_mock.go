package glass

import "github.com/gin-gonic/gin"

type ginEngineMockConnection struct {
}

func (inst *ginEngineMockConnection) init() EngineConnection {
	return inst
}

func (inst *ginEngineMockConnection) Close() error {
	return nil
}

func (inst *ginEngineMockConnection) Run() error {
	return nil
}

func (inst *ginEngineMockConnection) Join() error {
	return nil
}

func (inst *ginEngineMockConnection) RequestMapping(path string) EngineConnection {
	return inst
}

func (inst *ginEngineMockConnection) Handle(method string, path string, h gin.HandlerFunc) {
}

func (inst *ginEngineMockConnection) Filter(order int, h gin.HandlerFunc) {
}
