package factory

import (
	"net/http"
	"strconv"
	"time"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/bitwormhole/starter-restful/api/vo"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
)

type DefaultRestResponder struct {
	markup.Component `class:"glass-responder-registry"`
}

func (inst *DefaultRestResponder) _Impl() (glass.Responder, glass.ResponderRegistry) {
	return inst, inst
}

func (inst *DefaultRestResponder) GetRegistrationList() []*glass.ResponderRegistration {

	rr := &glass.ResponderRegistration{}
	rr.Name = "REST"
	rr.Responder = inst

	return []*glass.ResponderRegistration{rr}
}

func (inst *DefaultRestResponder) Send(resp *glass.Response) {
	if inst.hasError(resp) {
		inst.sendError(resp)
		return
	}
	inst.sendOk(resp)
}

func (inst *DefaultRestResponder) hasError(resp *glass.Response) bool {

	err := resp.Error
	if err != nil {
		return true
	}

	status := resp.Status
	if 0 < status && status < 200 {
		return true
	} else if 299 < status {
		return true
	}

	return false
}

func (inst *DefaultRestResponder) sendOk(resp *glass.Response) {

	o := resp.Data
	c := resp.Context
	status := resp.Status

	if status < 1 {
		status = http.StatusOK
	}

	o2, ok := o.(vo.ViewObject)
	if ok {
		meta := o2.GetMeta()
		inst.fillMeta(meta)
		meta.Status = status
	}

	c.JSON(status, o)
}

func (inst *DefaultRestResponder) sendError(resp *glass.Response) {

	o := &vo.BaseVO{}
	c := resp.Context
	err := resp.Error
	status := resp.Status

	if status < 1 {
		status = http.StatusInternalServerError
	}

	if err != nil {
		o.Error = err.Error()
	} else {
		o.Error = "Internal Server Error"
	}

	inst.fillMeta(&o.Meta)
	o.Status = status
	o.Message = "HTTP " + strconv.Itoa(status)
	c.JSON(status, o)
}

func (inst *DefaultRestResponder) fillMeta(meta *vo.Meta) {
	now := time.Now()
	meta.Time = util.NewTime(now)
	meta.TimeFormatted = now
}
