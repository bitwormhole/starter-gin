package containers

import (
	"github.com/bitwormhole/starter/application"
)

type GinLooper struct {
	Context application.Context
}

func (inst *GinLooper) Loop() error {
	pServer, err := inst.Context.FindComponent("#" + ID_GIN_WEB_SERVER)
	if err != nil {
		return err
	}
	container := pServer.(*GinServerContainer)
	return container.Run()
}
