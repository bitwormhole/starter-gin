// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package config_demo

import (
	glass0x47343f "github.com/bitwormhole/starter-gin/glass"
	demo0xa69cb5 "github.com/bitwormhole/starter-gin/src/demo/golang/demo"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComDemo1ctrl struct {
	instance *demo0xa69cb5.Demo1ctrl
	 markup0x23084a.Component `class:"rest-controller"`
	Resp glass0x47343f.MainResponder `inject:"#glass-main-responder"`
}


type pComDemo1ir struct {
	instance *demo0xa69cb5.Demo1ir
	 markup0x23084a.Component `class:"rest-interceptor-registry"`
}

