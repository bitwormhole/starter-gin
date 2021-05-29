package containers

import (
	"errors"
	"strconv"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
	"github.com/gin-gonic/gin"
)

const GinWebControllerClassName = "gin-web-controller"

type GinWebController interface {
	Config(engine *gin.Engine) error
}

type GinWebFilter interface {
	lang.PriorityProvider
	ConfigFilter(engine *gin.Engine) error
}

type GinServerContainer struct {
	port        int
	host        string
	engine      *gin.Engine
	filters     []GinWebFilter
	controllers []GinWebController
	appContext  application.Context
}

func (inst *GinServerContainer) Inject(context application.Context) error {

	ctrlist := []GinWebController{}
	filterlist := []GinWebFilter{}
	in := context.Injector()
	selector := "." + GinWebControllerClassName

	in.Inject(selector).AsList().To(func(o lang.Object) bool {
		controller, ok := o.(GinWebController)
		if ok {
			ctrlist = append(ctrlist, controller)
			return ok
		}
		filter, ok := o.(GinWebFilter)
		if ok {
			filterlist = append(filterlist, filter)
			return ok
		}
		return ok
	})

	inst.loadProperties(context)
	inst.controllers = ctrlist
	inst.filters = filterlist
	inst.appContext = context
	return in.Done()
}

func (inst *GinServerContainer) loadProperties(context application.Context) error {

	getter := context.GetProperties()
	str_host := getter.GetProperty("server.host", "")
	str_port := getter.GetProperty("server.port", "8080")

	port, err := strconv.ParseInt(str_port, 10, 32)
	if err != nil {
		return err
	}
	if port < 1 || 65535 < port {
		return errors.New("bad port value .")
	}

	inst.port = int(port)
	inst.host = str_host
	return nil
}

func (inst *GinServerContainer) initEngine() error {

	server := gin.Default()

	/*
		server.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"hello":  "world",
				"server": "gin",
				"port":   666,
			})
		})
	*/

	inst.engine = server
	return nil
}

func (inst *GinServerContainer) initFilters() error {

	filters := inst.filters
	engine := inst.engine

	for index := range filters {
		ctrl := filters[index]
		err := ctrl.ConfigFilter(engine)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *GinServerContainer) initControllers() error {

	controllers := inst.controllers
	engine := inst.engine

	for index := range controllers {
		ctrl := controllers[index]
		err := ctrl.Config(engine)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *GinServerContainer) initStatic() error {

	res := inst.appContext.GetResources()
	engine := inst.engine
	adapter := &GinFsAdapter{res: res}
	engine.StaticFS("/www", adapter.GetFS())
	return nil
}

func (inst *GinServerContainer) Run() error {

	port := inst.port
	host := inst.host

	if port < 1 {
		port = 8080
	}

	addr := host + ":" + strconv.Itoa(port)
	inst.engine.Run(addr)
	return nil
}

func (inst *GinServerContainer) Init() error {

	err := inst.initEngine()
	if err != nil {
		return err
	}

	err = inst.initFilters()
	if err != nil {
		return err
	}

	err = inst.initControllers()
	if err != nil {
		return err
	}

	err = inst.initStatic()
	if err != nil {
		return err
	}

	return nil
}

func (inst *GinServerContainer) Destroy() error {

	return nil
}
