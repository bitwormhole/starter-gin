package etc

import (
	"github.com/bitwormhole/starter-gin/demo/elements"
	"github.com/bitwormhole/starter-gin/security"
	"github.com/bitwormhole/starter-gin/web/filters"
	"github.com/bitwormhole/starter-gin/web/rest"
	"github.com/bitwormhole/starter/application"
	lang "github.com/bitwormhole/starter/lang"
)

func tokenSessionFilter(inst *filters.SecuritySessionFilter, ctx application.Context) error {

	// [component]
	//    class= gin-web-controller

	inst.SetSessionFactory(&security.TokenSessionFactory{})
	return nil
}

func theGinHandlerExample1(inst *elements.ExampleGinController, ctx application.Context) error {

	// [component]
	//    class= gin-web-controller

	in := ctx.Injector()

	in.Inject(".rest-context").To(func(o lang.Object) bool {
		rc, ok := o.(rest.Context)
		if ok {
			inst.REST = rc
		}
		return ok
	})

	return in.Done()
}
