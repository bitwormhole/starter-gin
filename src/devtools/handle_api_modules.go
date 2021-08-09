package devtools

import (
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/gin-gonic/gin"
)

type apiModulesHandler struct {
	modules []*ModuleInfoDTO
}

func (inst *apiModulesHandler) init(app application.Context) {

	const prefix = "module."
	const suffix = ".version"

	modlist := make([]*ModuleInfoDTO, 0)
	namelist := make([]string, 0)
	table := app.GetProperties().Export(nil)

	for key := range table {
		if strings.HasPrefix(key, prefix) && strings.HasSuffix(key, suffix) {
			name := key[len(prefix) : len(key)-len(suffix)]
			namelist = append(namelist, name)
		}
	}

	for _, name := range namelist {
		selector := "module." + name + "."
		mod := &ModuleInfoDTO{}
		mod.Index = table[selector+"index"]
		mod.Name = table[selector+"name"]
		mod.Version = table[selector+"version"]
		mod.Revision = table[selector+"revision"]
		if mod.Index == "" {
			continue
		}
		modlist = append(modlist, mod)
	}

	inst.modules = modlist
}

func (inst *apiModulesHandler) handle(c *gin.Context) {
	h := gin.H{
		"modules": inst.modules,
	}
	c.JSON(200, h)
}
