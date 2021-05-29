package etc

import (
	"github.com/bitwormhole/starter-gin/web/containers"
	"github.com/bitwormhole/starter-gin/web/rest"
	"github.com/bitwormhole/starter/application"
)

func theGinLooper(inst *containers.GinLooper, context application.Context) error {

	//	[component]
	// class=looper

	inst.Context = context
	return nil
}

func theGinServerContainer(inst *containers.GinServerContainer, context application.Context) error {

	//	[component]
	//  id=gin-web-server
	//  initMethod=Init
	// destroyMethod=Destroy

	inst.Inject(context)

	return nil
}

func theRestContext(inst *rest.SimpleRestContext, context application.Context) error {

	// [component]
	//	class=rest-context

	in := context.Injector()
	path := in.GetPropertyString("rest.context.path", "/api/v1")
	inst.ContextPath = path
	return in.Done()
}
