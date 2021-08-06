package glass

import (
	"sort"

	"github.com/gin-gonic/gin"
)

type innerFilterRegistration struct {
	order int
	fn    gin.HandlerFunc
}

////////////////////////////////////////////////////////////////////////////////

type innerFilterRegistrationSorter struct {
	items []*innerFilterRegistration
}

func (inst *innerFilterRegistrationSorter) sort(items []*innerFilterRegistration) {
	inst.items = items
	sort.Sort(inst)
}

func (inst *innerFilterRegistrationSorter) Swap(a int, b int) {
	item1 := inst.items[a]
	item2 := inst.items[b]
	inst.items[a] = item2
	inst.items[b] = item1
}

func (inst *innerFilterRegistrationSorter) Less(a int, b int) bool {
	item1 := inst.items[a]
	item2 := inst.items[b]
	return item1.order < item2.order
}

func (inst *innerFilterRegistrationSorter) Len() int {
	return len(inst.items)
}

////////////////////////////////////////////////////////////////////////////////
