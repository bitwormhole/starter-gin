package factory

import (
	"sort"
	"strings"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/gin-gonic/gin"
)

////////////////////////////////////////////////////////////////////////////////

type ginEngineConnectionHolder struct {
	handlers    []*HandlerRegistration
	filters     []*FilterRegistration
	hNoMethod   []*HandlerRegistration
	hNoResource []*HandlerRegistration
}

func (inst *ginEngineConnectionHolder) getEngineConnection(ctx glass.WebContext) glass.EngineConnection {
	ec := &ginEngineConnectionFacade{}
	ec.context = ctx
	ec.core = inst
	ec.path = "/"
	return ec
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
	inst.core.handlers = append(inst.core.handlers, r)
}

// HandleNoMethod 注册 error 处理器
func (inst *ginEngineConnectionFacade) HandleNoMethod(handler gin.HandlerFunc) {
	r := &HandlerRegistration{}
	r.Method = "error"
	r.Path = "/no/method"
	r.Handler = handler
	inst.core.hNoMethod = append(inst.core.hNoMethod, r)
}

// HandleNoResource 注册 error 处理器
func (inst *ginEngineConnectionFacade) HandleNoResource(handler gin.HandlerFunc) {
	r := &HandlerRegistration{}
	r.Method = "error"
	r.Path = "/no/resource"
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
	Method  string
	Path    string
	Handler gin.HandlerFunc
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
