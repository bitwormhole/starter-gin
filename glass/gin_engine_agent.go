package glass

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/bitwormhole/starter/util"
	"github.com/gin-gonic/gin"
)

type ginEngineRunnerFunc func(engine *gin.Engine) error

////////////////////////////////////////////////////////////////////////////////

type ginEngineAgent struct {
	engine *gin.Engine
	runner ginEngineRunnerFunc

	started  bool
	starting bool
	stopping bool
	stopped  bool

	filters  []*innerFilterRegistration
	handlers []*innerHandlerRegistration
}

func (inst *ginEngineAgent) openMock() (EngineConnection, error) {
	mock := &ginEngineMockConnection{}
	return mock.init(), nil
}

func (inst *ginEngineAgent) open() (EngineConnection, error) {
	conn := inst.init()
	return conn, nil
}

func (inst *ginEngineAgent) close() error {
	inst.stopping = true
	return nil
}

func (inst *ginEngineAgent) run() error {

	err := inst.applyFilters()
	if err != nil {
		return err
	}

	err = inst.applyHandlers()
	if err != nil {
		return err
	}

	defer func() { inst.stopped = true }()

	defer func() {
		x := recover()
		inst.handleRecover(x)
	}()

	inst.started = true

	return inst.runner(inst.engine)
}

func (inst *ginEngineAgent) applyFilters() error {
	engine := inst.engine
	list := inst.filters

	// TODO : sort

	for _, item := range list {
		engine.Use(item.fn)
	}
	return nil
}

func (inst *ginEngineAgent) applyHandlers() error {
	engine := inst.engine
	list := inst.handlers

	// TODO : sort

	for _, item := range list {
		method := item.method
		path := item.path
		engine.Handle(method, path, item.fn)
	}
	return nil
}

func (inst *ginEngineAgent) join() error {

	var interval time.Duration = 100 // ms
	var timeout time.Duration = 5000 // ms
	const unit = time.Millisecond

	for {
		if inst.started {
			break
		}
		if timeout > 0 {
			time.Sleep(unit * interval)
			timeout -= interval
		} else {
			return errors.New("timeout")
		}
	}

	interval = 1000
	timeout = -1

	for {
		time.Sleep(interval * unit)
		if inst.stopped {
			break
		}
	}

	return nil
}

func (inst *ginEngineAgent) handleRecover(x interface{}) {
	log.Println("revover from panic", x)
}

func (inst *ginEngineAgent) init() EngineConnection {

	inst.engine = gin.Default()
	inst.filters = make([]*innerFilterRegistration, 0)
	inst.handlers = make([]*innerHandlerRegistration, 0)

	conn := &ginEngineConnection{}
	conn.agent = inst
	conn.currentPath = "/"
	return conn
}

////////////////////////////////////////////////////////////////////////////////

type ginEngineConnection struct {
	agent       *ginEngineAgent
	currentPath string
}

func (inst *ginEngineConnection) _Impl() EngineConnection {
	return inst
}

func (inst *ginEngineConnection) Join() error {
	return inst.agent.join()
}

func (inst *ginEngineConnection) Run() error {
	return inst.agent.run()
}

func (inst *ginEngineConnection) Close() error {
	return inst.agent.close()
}

func (inst *ginEngineConnection) computePath(path string) string {
	if strings.HasPrefix(path, "/") {
		return path
	}
	return inst.currentPath + "/" + path
}

func (inst *ginEngineConnection) normalizePath(path string) string {
	builder := &util.PathBuilder{}
	builder.AppendPath(path)
	path = builder.String()
	return "/" + path
}

func (inst *ginEngineConnection) RequestMapping(path string) EngineConnection {

	path = inst.computePath(path)

	child := &ginEngineConnection{}
	child.agent = inst.agent
	child.currentPath = path
	return child
}

func (inst *ginEngineConnection) Filter(order int, h gin.HandlerFunc) {

	reg := &innerFilterRegistration{}
	reg.order = order
	reg.fn = h

	agent := inst.agent
	agent.filters = append(agent.filters, reg)
}

func (inst *ginEngineConnection) Handle(method string, path string, h gin.HandlerFunc) {

	path = inst.computePath(path)
	path = inst.normalizePath(path)

	reg := &innerHandlerRegistration{}
	reg.method = method
	reg.path = path
	reg.fn = h

	agent := inst.agent
	agent.handlers = append(agent.handlers, reg)
}

////////////////////////////////////////////////////////////////////////////////
