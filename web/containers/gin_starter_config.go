package containers

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/application/config"
	"github.com/bitwormhole/starter/lang"
)

const ID_GIN_WEB_SERVER = "gin-web-server"

func Config(cfg application.ConfigBuilder) {

	cfg.AddComponent(&config.ComInfo{

		ID: ID_GIN_WEB_SERVER,

		OnNew: func() lang.Object {
			return &GinServerContainer{}
		},

		OnInject: func(t lang.Object, context application.Context) error {
			return t.(*GinServerContainer).Inject(context)
		},

		OnInit: func(t lang.Object) error {
			return t.(*GinServerContainer).Init()
		},

		OnDestroy: func(t lang.Object) error {
			return t.(*GinServerContainer).Destroy()
		},
	})

}
