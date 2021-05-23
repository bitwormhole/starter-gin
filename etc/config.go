package etc

import (
	"github.com/bitwormhole/starter-gin/web/gin_starter"
	"github.com/bitwormhole/starter-gin/web/rest"
	"github.com/bitwormhole/starter/application"
)

func theGinLooper(inst *gin_starter.GinLooper, context application.Context) error {

	//	[component]
	// class=looper

	inst.Context = context
	return nil
}

func theGinServerContainer(inst *gin_starter.GinServerContainer, context application.Context) error {

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
