package glass

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
