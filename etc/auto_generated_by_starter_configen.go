// 这个文件是由 starter-configen 工具生成的配置代码，千万不要手工修改里面的任何内容。
package etc

import(
	gin_starter_b813dae6 "github.com/bitwormhole/starter-gin/web/gin_starter"
	rest_599fad3c "github.com/bitwormhole/starter-gin/web/rest"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)

func Config(cb application.ConfigBuilder) error {

    // theGinLooper
    cb.AddComponent(&config.ComInfo{
		ID: "theGinLooper",
		Class: "looper",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &gin_starter_b813dae6.GinLooper{}
		},
		OnInject: func(obj lang.Object,context application.Context) error {
		    target := obj.(*gin_starter_b813dae6.GinLooper)
		    return theGinLooper(target,context)
		},
    })

    // theGinServerContainer
    cb.AddComponent(&config.ComInfo{
		ID: "gin-web-server",
		Class: "",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &gin_starter_b813dae6.GinServerContainer{}
		},
		OnInit: func(obj lang.Object) error {
		    target := obj.(*gin_starter_b813dae6.GinServerContainer)
		    return target.Init()
		},
		OnDestroy: func(obj lang.Object) error {
		    target := obj.(*gin_starter_b813dae6.GinServerContainer)
		    return target.Destroy()
		},
		OnInject: func(obj lang.Object,context application.Context) error {
		    target := obj.(*gin_starter_b813dae6.GinServerContainer)
		    return theGinServerContainer(target,context)
		},
    })

    // theRestContext
    cb.AddComponent(&config.ComInfo{
		ID: "theRestContext",
		Class: "rest-context",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &rest_599fad3c.SimpleRestContext{}
		},
		OnInject: func(obj lang.Object,context application.Context) error {
		    target := obj.(*rest_599fad3c.SimpleRestContext)
		    return theRestContext(target,context)
		},
    })

    return nil
}

