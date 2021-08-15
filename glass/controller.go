package glass

// Controller 接口表示一个gin-glass控制器
type Controller interface {
	Init(ec EngineConnection) error
}

// ControllerFunc 提供一个面向函数的控制器接口
type ControllerFunc func(ec EngineConnection) error

////////////////////////////////////////////////////////////////////////////////

// GroupController 将一组 ControllerFunc 封装成一个 Controller 对象
type GroupController struct {
	fnlist  []ControllerFunc
	ctrlist []Controller
}

func (inst *GroupController) _Impl() Controller {
	return inst
}

// Init 初始化
func (inst *GroupController) Init(ec EngineConnection) error {

	list1 := inst.fnlist
	if list1 != nil {
		for _, fn := range list1 {
			ec2 := ec.RequestMapping(".")
			err := fn(ec2)
			if err != nil {
				return err
			}
		}
	}

	list2 := inst.ctrlist
	if list2 != nil {
		for _, ctrl := range list2 {
			ec2 := ec.RequestMapping(".")
			err := ctrl.Init(ec2)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// AddController 添加一个 Controller 到group
func (inst *GroupController) AddController(ctr Controller) {
	if ctr == nil {
		return
	}
	list := inst.ctrlist
	if list == nil {
		list = make([]Controller, 0)
	}
	list = append(list, ctr)
	inst.ctrlist = list
}

// AddFunction 添加一个func到group
func (inst *GroupController) AddFunction(fn ControllerFunc) {
	if fn == nil {
		return
	}
	list := inst.fnlist
	if list == nil {
		list = make([]ControllerFunc, 0)
	}
	list = append(list, fn)
	inst.fnlist = list
}
