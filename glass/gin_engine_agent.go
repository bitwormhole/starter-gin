package glass

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ginEngineRunnerFunc func(engine *gin.Engine) error

type ginEngineAgent struct {
	engine *gin.Engine
	runner ginEngineRunnerFunc

	started  bool
	starting bool
	stopping bool
	stopped  bool
}

func (inst *ginEngineAgent) openMock() (EngineConnection, error) {
	mock := &ginEngineMockConnection{}
	return mock.init(), nil
}

func (inst *ginEngineAgent) open() (EngineConnection, error) {
	inst.engine = gin.Default()
	conn := inst.myConnection()
	return conn, nil
}

func (inst *ginEngineAgent) close() error {
	inst.stopping = true
	return nil
}

func (inst *ginEngineAgent) run() error {

	defer func() { inst.stopped = true }()

	defer func() {
		x := recover()
		inst.handleRecover(x)
	}()

	inst.started = true

	return inst.runner(inst.engine)
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

func (inst *ginEngineAgent) myConnection() EngineConnection {
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

func (inst *ginEngineConnection) RequestMapping(path string) EngineConnection {

	if strings.HasPrefix(path, "/") {
		// NOP
	} else {
		path = inst.currentPath + "/" + path
	}

	child := &ginEngineConnection{}
	child.agent = inst.agent
	child.currentPath = path
	return child
}
