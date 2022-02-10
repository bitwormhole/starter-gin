package factory

import (
	"sort"
	"strconv"
	"strings"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/vlog"
	"github.com/gin-gonic/gin"
)

////////////////////////////////////////////////////////////////////////////////

type ginEngineConnectionHolder struct {
	handlers []*HandlerRegistration
	filters  []*FilterRegistration
	// interceptors []*glass.InterceptorRegistration
	hNoMethod   []*HandlerRegistration
	hNoResource []*HandlerRegistration
}

func (inst *ginEngineConnectionHolder) getEngineConnection(ctx glass.WebContext) glass.EngineConnection {

	// inst.interceptors = inst.getInterceptorRegistrations(ctx.GetInterceptors())

	ec := &ginEngineConnectionFacade{}
	ec.context = ctx
	ec.core = inst
	ec.path = "/"
	return ec
}

func (inst *ginEngineConnectionHolder) getInterceptorRegistrations(src []glass.InterceptorRegistry) []*glass.InterceptorRegistration {
	dst := make([]*glass.InterceptorRegistration, 0)
	for _, item1 := range src {
		list := item1.GetRegistrationList()
		for _, item2 := range list {
			dst = append(dst, item2)
		}
	}
	return dst
}

// （如果需要） 为所有的 handler 逐个创建拦截器
func (inst *ginEngineConnectionHolder) prepareInterceptorForHandlers(icbTable map[string]*interceptorChainBuilder) {
	handlers := inst.handlers
	for _, item := range handlers {
		icb := icbTable[item.ContextID]
		if icb == nil {
			continue
		}
		h := item.Handler
		item.Handler = icb.makeChainFor(h)
	}
}

func (inst *ginEngineConnectionHolder) applyToEngine(e *gin.Engine) error {

	err := inst.configureFilters(e)
	if err != nil {
		return err
	}

	err = inst.configureHandlers(e)
	if err != nil {
		return err
	}

	err = inst.configureNoMethod(e)
	if err != nil {
		return err
	}

	err = inst.configureNoRes(e)
	if err != nil {
		return err
	}

	return nil
}

func (inst *ginEngineConnectionHolder) configureFilters(e *gin.Engine) error {
	src := inst.filters
	dst := make([]gin.HandlerFunc, 0)
	(&filterRegistrationSort{list: src}).sort()
	for _, item := range src {
		if item == nil {
			continue
		}
		h := item.Filter
		if h == nil {
			continue
		}
		dst = append(dst, h)
		hstr := lang.StringifyObject(h)
		vlog.Info("gin.Engine.Use(middleware:", hstr, ")")
	}
	e.Use(dst...)
	return nil
}

func (inst *ginEngineConnectionHolder) configureHandlers(e *gin.Engine) error {
	src := inst.handlers
	(&handlerRegistrationSort{list: src}).sort()
	for _, item := range src {
		if item == nil {
			continue
		}
		h := item.Handler
		if h == nil {
			continue
		}
		e.Handle(item.Method, item.Path, h)
	}
	return nil
}

func (inst *ginEngineConnectionHolder) configureNoMethod(e *gin.Engine) error {
	dst := make([]gin.HandlerFunc, 0)
	all := inst.hNoMethod
	(&handlerRegistrationSort{list: all}).sort()
	for _, item := range all {
		if item == nil {
			continue
		}
		h := item.Handler
		if h == nil {
			continue
		}
		dst = append(dst, h)
	}
	if len(dst) > 0 {
		e.NoMethod(dst...)
	}
	return nil
}

func (inst *ginEngineConnectionHolder) configureNoRes(e *gin.Engine) error {
	dst := make([]gin.HandlerFunc, 0)
	all := inst.hNoResource
	(&handlerRegistrationSort{list: all}).sort()
	for _, item := range all {
		if item == nil {
			continue
		}
		h := item.Handler
		if h == nil {
			continue
		}
		dst = append(dst, h)
	}
	if len(dst) > 0 {
		e.NoRoute(dst...)
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type ginEngineConnectionFacade struct {
	core    *ginEngineConnectionHolder
	context glass.WebContext
	path    string
}

func (inst *ginEngineConnectionFacade) _Impl() glass.EngineConnection {
	return inst
}

func (inst *ginEngineConnectionFacade) normalizePath(path string) string {
	const slash = "/"
	endsWithSlash := strings.HasSuffix(path, slash)
	builder := strings.Builder{}
	items := strings.Split(path, slash)
	for _, it := range items {
		it = strings.TrimSpace(it)
		if it == "" {
			continue
		}
		builder.WriteString(slash)
		builder.WriteString(it)
	}
	if endsWithSlash {
		builder.WriteString(slash)
	}
	return builder.String()
}

func (inst *ginEngineConnectionFacade) href(path string) string {
	if strings.HasPrefix(path, "/") {
		return inst.normalizePath(path)
	}
	return inst.normalizePath(inst.path + "/" + path)
}

// GetWebContext 取 WebContext
func (inst *ginEngineConnectionFacade) GetWebContext() glass.WebContext {
	return inst.context
}

// RequestMapping 映射到新的路径
func (inst *ginEngineConnectionFacade) RequestMapping(path string) glass.EngineConnection {

	path2 := inst.href(path)
	ec2 := &ginEngineConnectionFacade{}

	ec2.context = inst.context
	ec2.core = inst.core
	ec2.path = path2

	return ec2
}

// Filter 注册过滤器
func (inst *ginEngineConnectionFacade) Filter(order int, handler gin.HandlerFunc) {
	r := &FilterRegistration{}
	r.Order = order
	r.Filter = handler
	inst.core.filters = append(inst.core.filters, r)
}

// Handle 注册处理器
func (inst *ginEngineConnectionFacade) Handle(method string, path string, handler gin.HandlerFunc) {
	r := &HandlerRegistration{}
	r.Method = method
	r.Path = inst.href(path)
	r.Handler = handler
	r.ContextID = inst.context.GetContextID()
	inst.core.handlers = append(inst.core.handlers, r)
}

// HandleNoMethod 注册 error 处理器
func (inst *ginEngineConnectionFacade) HandleNoMethod(order int, handler gin.HandlerFunc) {
	r := &HandlerRegistration{}
	r.Method = "error"
	r.Path = "/no/method/" + strconv.Itoa(order+10000000)
	r.Handler = handler
	inst.core.hNoMethod = append(inst.core.hNoMethod, r)
}

// HandleNoResource 注册 error 处理器
func (inst *ginEngineConnectionFacade) HandleNoResource(order int, handler gin.HandlerFunc) {
	r := &HandlerRegistration{}
	r.Method = "error"
	r.Path = "/no/resource/" + strconv.Itoa(order+10000000)
	r.Handler = handler
	inst.core.hNoResource = append(inst.core.hNoResource, r)
}

////////////////////////////////////////////////////////////////////////////////

// FilterRegistration ...
type FilterRegistration struct {
	Order  int
	Filter gin.HandlerFunc
}

type filterRegistrationSort struct {
	list []*FilterRegistration
}

func (inst *filterRegistrationSort) sort() {
	sort.Sort(inst)
}

func (inst *filterRegistrationSort) Len() int {
	if inst.list == nil {
		return 0
	}
	return len(inst.list)
}

func (inst *filterRegistrationSort) Less(a, b int) bool {
	list := inst.list
	return list[a].Order < list[b].Order
}

func (inst *filterRegistrationSort) Swap(a, b int) {
	list := inst.list
	aa := list[a]
	bb := list[b]
	list[a] = bb
	list[b] = aa
}

////////////////////////////////////////////////////////////////////////////////

// HandlerRegistration ...
type HandlerRegistration struct {
	Method    string
	Path      string
	ContextID string
	Handler   gin.HandlerFunc
}

type handlerRegistrationSort struct {
	list []*HandlerRegistration
}

func (inst *handlerRegistrationSort) sort() {
	sort.Sort(inst)
}

func (inst *handlerRegistrationSort) Len() int {
	if inst.list == nil {
		return 0
	}
	return len(inst.list)
}

func (inst *handlerRegistrationSort) Less(a, b int) bool {
	list := inst.list
	s1 := list[a].Path + "#" + list[a].Method
	s2 := list[b].Path + "#" + list[b].Method
	return strings.Compare(s1, s2) < 0
}

func (inst *handlerRegistrationSort) Swap(a, b int) {
	list := inst.list
	aa := list[a]
	bb := list[b]
	list[a] = bb
	list[b] = aa
}

////////////////////////////////////////////////////////////////////////////////
