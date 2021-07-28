package web

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bitwormhole/starter/application"
	"github.com/gin-gonic/gin"
)

type GinStarter struct {
	runtime *myGinRuntime

	ContextPath        string
	Host               string
	Port               int
	ApplicationContext application.Context
}

func (inst *GinStarter) __impl__() application.Looper {
	return inst
}

func (inst *GinStarter) Start() error {

	rt := inst.runtime

	if rt == nil {
		rt = &myGinRuntime{}
	} else {
		if rt.isStopped() {
			rt = &myGinRuntime{}
		} else {
			return nil
		}
	}

	rt.config = inst
	inst.runtime = rt
	return rt.start()
}

func (inst *GinStarter) Stop() error {
	rt := inst.runtime
	if rt == nil {
		return nil
	}
	return rt.stop()
}

func (inst *GinStarter) Loop() error {
	rt := inst.runtime
	if rt == nil {
		return nil
	}
	return rt.waitForStopped()
}

////////////////////////////////////////////////////////////////////////////////

type myGinRuntime struct {
	engine *gin.Engine
	config *GinStarter

	endpoint  string // host:port
	LastError error

	started  bool
	stopped  bool
	starting bool
	stopping bool
}

func (inst *myGinRuntime) stop() error {
	inst.stopping = true
	return nil
}

func (inst *myGinRuntime) isStopped() bool {
	return inst.stopped
}

func (inst *myGinRuntime) loop() error {

	engine := inst.engine

	return engine.Run(inst.endpoint)
}

func (inst *myGinRuntime) start() error {
	inst.starting = true
	go func() { inst.LastError = inst.run() }()
	return inst.waitForReady()
}

func (inst *myGinRuntime) waitForReady() error {
	for {
		if inst.started {
			break
		}

		if inst.stopped {
			return inst.LastError
		}

		time.Sleep(time.Duration(100 * 1000 * 1000))
	}
	return nil
}

func (inst *myGinRuntime) waitForStopped() error {
	for {
		if inst.stopped {
			return inst.LastError
		}
		time.Sleep(time.Duration(1000 * 1000 * 1000))
	}
}

func (inst *myGinRuntime) run() error {

	defer func() { inst.stopped = true }()

	defer func() {
		err := recover()
		if err == nil {
			return
		}
		inst.LastError = errors.New(fmt.Sprint(err))
		fmt.Println("recover error:", err)
	}()

	err := inst.init()
	if err != nil {
		return err
	}
	inst.started = true
	return inst.loop()
}

func (inst *myGinRuntime) init() error {

	facade := &myGinStarterFacade{}
	starter := &myGinStarterCore{}

	facade.core = starter
	facade.contextPath = inst.config.ContextPath

	starter.runtime = inst
	starter.facade = facade
	starter.controllers = make([]Controller, 0)
	starter.filters = make([]*myFilterReg, 0)
	starter.handlers = make([]*myHandlerReg, 0)

	return starter.init()
}

////////////////////////////////////////////////////////////////////////////////

type myGinStarterCore struct {
	runtime *myGinRuntime
	facade  *myGinStarterFacade

	handlerOnNoMethod gin.HandlerFunc
	handlerOnNoRes    gin.HandlerFunc

	controllers []Controller
	filters     []*myFilterReg
	handlers    []*myHandlerReg
}

func (inst *myGinStarterCore) init() error {

	inst.runtime.engine = gin.Default()

	err := inst.init_endpoint()
	if err != nil {
		return err
	}

	err = inst.init_controllers_1()
	if err != nil {
		return err
	}

	err = inst.init_controllers_2()
	if err != nil {
		return err
	}

	err = inst.init_filters()
	if err != nil {
		return err
	}

	err = inst.init_handlers()
	if err != nil {
		return err
	}

	err = inst.init_error_handlers()
	if err != nil {
		return err
	}

	return nil
}

func (inst *myGinStarterCore) init_endpoint() error {

	runtime := inst.runtime
	port := runtime.config.Port
	host := runtime.config.Host

	if port < 1 {
		port = 8080
	}

	fmt.Println("config gin.server.host=", host)
	fmt.Println("config gin.server.port=", port)

	runtime.endpoint = host + ":" + strconv.Itoa(port)
	return nil
}

func (inst *myGinStarterCore) init_controllers_1() error {

	app := inst.runtime.config.ApplicationContext
	list1, err := app.GetComponentList("*")
	list2 := make([]Controller, 0)

	if err != nil {
		return err
	}

	for index := range list1 {
		obj := list1[index]
		ctrl, ok := obj.(Controller)
		if ok {
			list2 = append(list2, ctrl)
		}
	}

	inst.controllers = list2
	return nil
}

func (inst *myGinStarterCore) init_controllers_2() error {

	list := inst.controllers
	container := inst.facade

	for index := range list {
		ctrl := list[index]
		err := ctrl.Config(container)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myGinStarterCore) init_filters() error {

	engine := inst.runtime.engine
	list := inst.filters
	sort.Sort(&myFilterRegSorter{items: list})

	for index := range list {
		item := list[index]
		engine.Use(item.fn)
	}

	return nil
}

func (inst *myGinStarterCore) init_handlers() error {

	engine := inst.runtime.engine
	list := inst.handlers
	sort.Sort(&myHandlerRegSorter{items: list})

	for index := range list {
		item := list[index]
		engine.Handle(item.GetMethod(), item.GetPath(), item.fn)
	}

	return nil
}

func (inst *myGinStarterCore) init_error_handlers() error {

	engine := inst.runtime.engine
	no_method := inst.handlerOnNoMethod
	no_res := inst.handlerOnNoRes

	if no_method != nil {
		engine.NoMethod(no_method)
	}

	if no_res != nil {
		engine.NoRoute(no_res)
	}

	return nil
}

func (inst *myGinStarterCore) add_handler(reg *myHandlerReg) {
	inst.handlers = append(inst.handlers, reg)
}

func (inst *myGinStarterCore) add_filter(reg *myFilterReg) {
	inst.filters = append(inst.filters, reg)
}

////////////////////////////////////////////////////////////////////////////////

type myGinStarterFacade struct {
	core        *myGinStarterCore
	contextPath string
}

func (inst *myGinStarterFacade) Engine() *gin.Engine {
	return inst.core.runtime.engine
}

func (inst *myGinStarterFacade) Mapping(path string) Container {

	child := &myGinStarterFacade{}
	child.core = inst.core

	if strings.HasPrefix(path, "/") {
		child.contextPath = path
	} else {
		child.contextPath = inst.contextPath + "/" + path
	}

	return child
}

func (inst *myGinStarterFacade) HandleNoMethod(fn gin.HandlerFunc) {
	inst.core.handlerOnNoMethod = fn
}

func (inst *myGinStarterFacade) HandleNoResource(fn gin.HandlerFunc) {
	inst.core.handlerOnNoRes = fn
}

func (inst *myGinStarterFacade) AddFilter() FilterRegistration {
	reg := &myFilterReg{}
	reg.core = inst.core
	reg.priority = 0
	return reg
}

func (inst *myGinStarterFacade) GET(path string) HandlerRegistration {
	return inst.Handle("GET", path)
}

func (inst *myGinStarterFacade) POST(path string) HandlerRegistration {
	return inst.Handle("POST", path)
}

func (inst *myGinStarterFacade) PUT(path string) HandlerRegistration {
	return inst.Handle("PUT", path)
}

func (inst *myGinStarterFacade) DELETE(path string) HandlerRegistration {
	return inst.Handle("DELETE", path)
}

func (inst *myGinStarterFacade) Handle(method string, path string) HandlerRegistration {
	reg := &myHandlerReg{}
	reg.core = inst.core
	reg.method = method
	reg.path1 = inst.contextPath
	reg.path2 = path
	return reg
}

////////////////////////////////////////////////////////////////////////////////

type myFilterReg struct {
	core     *myGinStarterCore
	fn       gin.HandlerFunc
	priority int
}

func (inst *myFilterReg) Priority(priority int) FilterRegistration {
	inst.priority = priority
	return inst
}

func (inst *myFilterReg) Container() Container {
	return inst.core.facade
}

func (inst *myFilterReg) Handle(h gin.HandlerFunc) {
	if h == nil {
		return
	}
	inst.fn = h
	inst.core.add_filter(inst)
}

/////////////////////////////////////////////

type myFilterRegSorter struct {
	items []*myFilterReg
}

func (inst *myFilterRegSorter) Len() int {
	return len(inst.items)
}

func (inst *myFilterRegSorter) Less(a, b int) bool {
	aa := inst.items[a].priority
	bb := inst.items[b].priority
	return aa < bb
}

func (inst *myFilterRegSorter) Swap(a, b int) {
	aa := inst.items[a]
	bb := inst.items[b]
	inst.items[a] = bb
	inst.items[b] = aa
}

////////////////////////////////////////////////////////////////////////////////

type myHandlerReg struct {
	core      *myGinStarterCore
	fn        gin.HandlerFunc
	method    string
	path1     string
	path2     string
	pathFinal string
}

func (inst *myHandlerReg) Method(method string) HandlerRegistration {
	inst.method = method
	return inst
}

func (inst *myHandlerReg) Mapping(path string) HandlerRegistration {
	inst.path2 = path
	return inst
}

func (inst *myHandlerReg) Container() Container {
	return inst.core.facade
}

func (inst *myHandlerReg) Handle(h gin.HandlerFunc) {
	if h == nil {
		return
	}
	inst.fn = h
	inst.core.add_handler(inst)
}

func (inst *myHandlerReg) GetMethod() string {
	method := inst.method
	if method == "" {
		method = "GET"
	}
	return strings.ToUpper(method)
}

func (inst *myHandlerReg) GetPath() string {
	path := inst.pathFinal
	if path == "" {
		path = inst.computeFinalPath()
		inst.pathFinal = path
	}
	return path
}

func (inst *myHandlerReg) computeFinalPath() string {

	p1 := inst.path1
	p2 := inst.path2
	path := p1 + "/" + p2

	if strings.HasPrefix(p2, "/") {
		path = p2
	}

	list := strings.Split(path, "/")
	path = ""

	for index := range list {
		str := list[index]
		if str == "" {
			continue
		} else if str == "." {
			continue
		} else if str == ".." {
			continue
		} else {
			path = path + "/" + str
		}
	}

	if path == "" {
		return "/"
	}

	return path
}

/////////////////////////////////////////////

type myHandlerRegSorter struct {
	items []*myHandlerReg
}

func (inst *myHandlerRegSorter) Len() int {
	return len(inst.items)
}

func (inst *myHandlerRegSorter) Less(a, b int) bool {
	aa := inst.items[a].GetPath()
	bb := inst.items[b].GetPath()
	return strings.Compare(aa, bb) < 0
}

func (inst *myHandlerRegSorter) Swap(a, b int) {
	aa := inst.items[a]
	bb := inst.items[b]
	inst.items[a] = bb
	inst.items[b] = aa
}

////////////////////////////////////////////////////////////////////////////////
