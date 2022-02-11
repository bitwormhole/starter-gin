package factory

import (
	"errors"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter/markup"
)

// TheMainResponder ...
type TheMainResponder struct {
	markup.Component `id:"glass-main-responder" initMethod:"Init"`

	TargetName string                    `inject:"${glass.target.responder.name}"`
	Responders []glass.ResponderRegistry `inject:".glass-responder-registry"`

	all    map[string]*glass.ResponderRegistration
	target glass.Responder
}

func (inst *TheMainResponder) _Impl() glass.MainResponder {
	return inst
}

func (inst *TheMainResponder) Init() error {
	return inst.loadTarget()
}

func (inst *TheMainResponder) loadTarget() error {

	dst := make(map[string]*glass.ResponderRegistration)
	src := inst.Responders

	for _, item1 := range src {
		mid := item1.GetRegistrationList()
		for _, item2 := range mid {
			name := item2.Name
			dst[name] = item2
		}
	}

	name := inst.TargetName
	reg := dst[name]
	if reg == nil {
		return errors.New("no glass.Responder named " + name)
	}
	target := reg.Responder
	if target == nil {
		return errors.New("the glass.Responder is nil, named " + name)
	}

	inst.all = dst
	inst.target = target
	return nil
}

func (inst *TheMainResponder) TargetResponderName() string {
	return inst.TargetName
}

func (inst *TheMainResponder) Send(resp *glass.Response) {
	inst.target.Send(resp)
}
