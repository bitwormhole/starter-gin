package glass

// WebContext 表示一个web上下文(service)
type WebContext struct {
	Container   *Container
	ContextPath string
	Controllers []Controller
}

// MainController 取上下文的主控制器
func (inst *WebContext) MainController() Controller {
	return inst
}

// Init  控制器初始化
func (inst *WebContext) Init(ec EngineConnection) error {
	ec = ec.RequestMapping(inst.ContextPath)
	list := inst.Controllers
	if list == nil {
		return nil
	}
	for _, ctrl := range list {
		if ctrl == nil {
			continue
		}
		err := ctrl.Init(ec)
		if err != nil {
			return err
		}
	}
	return nil
}
