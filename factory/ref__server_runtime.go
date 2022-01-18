package glass

import (
	"strings"

	"github.com/bitwormhole/starter/lang"
)

type serverRuntime struct {
	Container *Container

	// private
	indexPageNames []string
	connections    []EngineConnection
	pool           lang.ReleasePool
}

func (inst *serverRuntime) start() error {
	worker := &serverRuntimeStarter{runtime: inst}
	return worker.start()
}

func (inst *serverRuntime) stop() error {
	worker := &serverRuntimeStopper{runtime: inst}
	return worker.stop()
}

func (inst *serverRuntime) loop() error {
	worker := &serverRuntimeLooper{runtime: inst}
	return worker.loop()
}

////////////////////////////////////////////////////////////////////////////////

type serverRuntimeStarter struct {
	runtime *serverRuntime
}

func (inst *serverRuntimeStarter) start() error {

	err := inst.initPool()
	if err != nil {
		return err
	}

	err = inst.initIndexPageNames()
	if err != nil {
		return err
	}

	err = inst.initConnections()
	if err != nil {
		return err
	}

	err = inst.initContexts()
	if err != nil {
		return err
	}

	return nil
}

func (inst *serverRuntimeStarter) initPool() error {
	rt := inst.runtime
	pool := lang.CreateReleasePool()
	rt.pool = pool
	return nil
}

func (inst *serverRuntimeStarter) initConnections() error {
	rt := inst.runtime
	connectors := rt.Container.Connectors
	connections := make([]EngineConnection, 0)
	for _, connector := range connectors {
		connection, err := connector.Open()
		if err != nil {
			return err
		}
		connections = append(connections, connection)
	}
	rt.connections = connections
	return nil
}

func (inst *serverRuntimeStarter) initIndexPageNames() error {

	rt := inst.runtime
	container := rt.Container
	text := container.IndexPages

	text = strings.ReplaceAll(text, " ", ",")
	array := strings.Split(text, ",")
	namelist := make([]string, 0)

	for _, item := range array {
		item = strings.TrimSpace(item)
		if item != "" {
			namelist = append(namelist, item)
		}
	}

	rt.indexPageNames = namelist
	return nil
}

func (inst *serverRuntimeStarter) initContexts() error {
	rt := inst.runtime
	container := rt.Container
	services := container.Services
	connections := rt.connections
	for _, service := range services {
		for _, conn := range connections {
			err := service.MainController().Init(conn)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type serverRuntimeStopper struct {
	runtime *serverRuntime
}

func (inst *serverRuntimeStopper) stop() error {
	pool := inst.runtime.pool
	return pool.Release()
}

////////////////////////////////////////////////////////////////////////////////

type serverRuntimeLooper struct {
	runtime *serverRuntime
}

func (inst *serverRuntimeLooper) loop() error {

	connections := inst.runtime.connections

	for _, conn := range connections {
		err := inst.startEngineLoop(conn)
		if err != nil {
			return err
		}
	}

	for _, conn := range connections {
		err := conn.Join()
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *serverRuntimeLooper) startEngineLoop(ec EngineConnection) error {
	go func() {
		err := ec.Run()
		inst.handleWorkerError(err)
	}()
	return nil
}

func (inst *serverRuntimeLooper) handleWorkerError(err error) {
	if err == nil {
		return
	}
	panic(err)
}

////////////////////////////////////////////////////////////////////////////////
