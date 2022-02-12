package factory

import (
	"net/http"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/gin-gonic/gin"
)

type ginEngineBuilder struct {
}

func (inst *ginEngineBuilder) makeMainHandler(c glass.Container) (http.Handler, error) {
	ecHolder := &ginEngineConnectionHolder{}
	ctxList := c.GetContexts()
	icbTable := make(map[string]*interceptorChainBuilder)

	for _, ctx := range ctxList {
		ec := ecHolder.getEngineConnection(ctx)
		err := inst.scanContext(ctx, ec)
		if err != nil {
			return nil, err
		}
		icbTable[ctx.GetContextID()] = inst.makeInterceptorChainBuilder(ctx)
	}

	ecHolder.prepareInterceptorForHandlers(icbTable)

	// engine := gin.New()
	engine := gin.Default()

	err := ecHolder.applyToEngine(engine)
	if err != nil {
		return nil, err
	}
	return engine, nil
}

func (inst *ginEngineBuilder) makeInterceptorChainBuilder(ctx glass.WebContext) *interceptorChainBuilder {
	icb := &interceptorChainBuilder{}
	icb.initWithRegistryList(ctx.GetInterceptors())
	return icb
}

func (inst *ginEngineBuilder) scanContext(ctx glass.WebContext, ec glass.EngineConnection) error {
	ctrList := ctx.GetControllers()
	for _, controller := range ctrList {
		contextPath := ctx.GetContextPath()
		ec2 := ec.RequestMapping(contextPath)
		err := controller.Init(ec2)
		if err != nil {
			return err
		}
	}

	// icb := interceptorChainBuilder{}
	// icb.initWithRegistryList(ctx.GetInterceptors())
	// icb.makeChainFor(nil)

	return nil
}

func (inst *ginEngineBuilder) makeNetworkConnList(c glass.Container) ([]glass.NetworkConnection, error) {
	h, err := inst.makeMainHandler(c)
	if err != nil {
		return nil, err
	}
	src := c.GetConnectors()
	dst := make([]glass.NetworkConnection, 0)
	for _, connector := range src {
		if !connector.Enabled() {
			continue
		}
		conn, err := connector.Connect(h)
		if err != nil {
			return nil, err
		}
		dst = append(dst, conn)
	}
	return dst, nil
}
