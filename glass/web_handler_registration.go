package glass

import (
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

type innerHandlerRegistration struct {
	method string
	path   string
	fn     gin.HandlerFunc
	isdir  bool
}

////////////////////////////////////////////////////////////////////////////////

type innerHandlerRegistrationSorter struct {
	items []*innerHandlerRegistration
}

func (inst *innerHandlerRegistrationSorter) sort(items []*innerHandlerRegistration) {
	inst.items = items
	sort.Sort(inst)
}

func (inst *innerHandlerRegistrationSorter) Swap(a int, b int) {
	item1 := inst.items[a]
	item2 := inst.items[b]
	inst.items[a] = item2
	inst.items[b] = item1
}

func (inst *innerHandlerRegistrationSorter) Less(a int, b int) bool {
	item1 := inst.items[a]
	item2 := inst.items[b]
	str1 := item1.path + "#" + item1.method
	str2 := item2.path + "#" + item2.method
	return strings.Compare(str1, str2) < 0
}

func (inst *innerHandlerRegistrationSorter) Len() int {
	return len(inst.items)
}

////////////////////////////////////////////////////////////////////////////////
