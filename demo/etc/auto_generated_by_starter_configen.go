// 这个文件是由 starter-configen 工具生成的配置代码，千万不要手工修改里面的任何内容。
package etc

import(
	elements_9f940d07 "github.com/bitwormhole/starter-gin/demo/elements"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)

func Config(cb application.ConfigBuilder) error {

    // theGinHandlerExample1
    cb.AddComponent(&config.ComInfo{
		ID: "theGinHandlerExample1",
		Class: "gin-web-controller",
		Scope: application.ScopeSingleton,
		Aliases: []string{},
		OnNew: func() lang.Object {
		    return &elements_9f940d07.ExampleGinController{}
		},
		OnInject: func(obj lang.Object,context application.Context) error {
		    target := obj.(*elements_9f940d07.ExampleGinController)
		    return theGinHandlerExample1(target,context)
		},
    })

    return nil
}

