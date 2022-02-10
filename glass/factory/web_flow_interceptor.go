package factory

import (
	"sort"

	"github.com/bitwormhole/starter-gin/glass"
	"github.com/gin-gonic/gin"
)

type interceptorChainBuilder struct {
	reglist []*glass.InterceptorRegistration
}

func (inst *interceptorChainBuilder) initWithRegistryList(src []glass.InterceptorRegistry) {
	dst := make([]*glass.InterceptorRegistration, 0)
	for _, item1 := range src {
		list := item1.GetRegistrationList()
		for _, item2 := range list {
			dst = append(dst, item2)
		}
	}
	inst.initWithRegistrationList(dst)
}

func (inst *interceptorChainBuilder) initWithRegistrationList(src []*glass.InterceptorRegistration) {
	dst := make([]*glass.InterceptorRegistration, 0)
	for _, item := range src {
		if item == nil {
			continue
		}
		if item.Interceptor == nil {
			continue
		}
		dst = append(dst, item)
	}
	inst.reglist = dst
	sort.Sort(inst)
}

func (inst *interceptorChainBuilder) Len() int {
	return len(inst.reglist)
}
func (inst *interceptorChainBuilder) Less(a, b int) bool {
	o1 := inst.reglist[a].Order
	o2 := inst.reglist[b].Order
	return o1 < o2
}
func (inst *interceptorChainBuilder) Swap(a, b int) {
	list := inst.reglist
	o1 := list[a]
	o2 := list[b]
	list[a] = o2
	list[b] = o1
}

func (inst *interceptorChainBuilder) makeChainFor(target gin.HandlerFunc) gin.HandlerFunc {
	h := target
	list := inst.reglist
	for i := len(list) - 1; i >= 0; i-- {
		item := list[i]
		h = item.Interceptor.Intercept(h)
	}
	return h
}

////////////////////////////////////////////////////////////////////////////////
