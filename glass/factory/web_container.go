package factory

import (
	"strings"
	"time"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

// WebContainer glass 的核心容器
type WebContainer struct {
	markup.Component `id:"gin-web-container"`

	// Server       glass.Server             `inject:"#gin-web-server"`

	Name         string                   `inject:"${server.name}"`
	IndexPages   string                   `inject:"${web.static.index-pages}"`
	ContentTypes glass.ContentTypeManager `inject:"#gin-web-content-types"`
	Connectors   []glass.Connector        `inject:".web-server-connector"`
	Contexts     []glass.WebContext       `inject:".web-context"`

	networks []glass.NetworkConnection
}

func (inst *WebContainer) _Impl() (glass.Container, application.LifeRegistry) {
	return inst, inst
}

// GetLifeRegistration 取生命周期注册项
func (inst *WebContainer) GetLifeRegistration() *application.LifeRegistration {
	r := &application.LifeRegistration{}
	r.OnInit = inst.onInit
	r.Looper = inst
	return r
}

// GetName 取容器名称
func (inst *WebContainer) GetName() string {
	return inst.Name
}

// GetIndexPageNames 取index页面名称列表
func (inst *WebContainer) GetIndexPageNames() []string {
	dst := make([]string, 0)
	src := strings.Split(inst.IndexPages, ",")
	for _, el := range src {
		el = strings.TrimSpace(el)
		if len(el) < 1 {
			continue
		}
		dst = append(dst, el)
	}
	return dst
}

// GetConnectors 取连接器列表
func (inst *WebContainer) GetConnectors() []glass.Connector {
	return inst.Connectors
}

// GetContexts 取WEB上下文列表
func (inst *WebContainer) GetContexts() []glass.WebContext {
	return inst.Contexts
}

// GetContentTypes 取文件类型管理器
func (inst *WebContainer) GetContentTypes() glass.ContentTypeManager {
	return inst.ContentTypes
}

// GetIndexPages 取index页面名称列表（in string）
func (inst *WebContainer) GetIndexPages() string {
	return inst.IndexPages
}

func (inst *WebContainer) onInit() error {
	builder := ginEngineBuilder{}
	nets, err := builder.makeNetworkConnList(inst)
	if err != nil {
		return err
	}
	inst.networks = nets
	return nil
}

// func (inst *WebContainer) onStart() error {
// 	return nil
// }

// Loop 运行网络循环
func (inst *WebContainer) Loop() error {
	src := inst.networks
	dst := make([]*networkRunner, 0)
	for _, nc := range src {
		runner := &networkRunner{nc: nc}
		dst = append(dst, runner)
		runner.start()
	}
	for _, runner := range dst {
		runner.join()
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type networkRunner struct {
	nc glass.NetworkConnection

	started bool
	stopped bool
}

func (inst *networkRunner) start() {
	go func() {
		inst.run()
	}()
	inst.started = true
}

func (inst *networkRunner) run() {

	defer func() {
		e := recover()
		inst.logError(e)
	}()

	defer func() {
		inst.stopped = true
	}()

	err := inst.nc.Run()
	inst.logError(err)
}

func (inst *networkRunner) logError(o lang.Object) {
	if o == nil {
		return
	}
	vlog.Error(o)
}

func (inst *networkRunner) join() {
	if !inst.started {
		return
	}
	for {
		if inst.stopped {
			break
		}
		time.Sleep(time.Second * 2)
	}
}
