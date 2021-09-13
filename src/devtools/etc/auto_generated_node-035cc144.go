// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package etc

import (
	devtools0x08d6b0 "github.com/bitwormhole/starter-gin/src/devtools"
	application0x67f6c5 "github.com/bitwormhole/starter/application"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComDevtoolsController struct {
	instance *devtools0x08d6b0.DevtoolsController
	 markup0x23084a.Component `class:"rest-controller"`
	AppContext application0x67f6c5.Context `inject:"context"`
	Enable bool `inject:"${web.devtools.enable}"`
	accumulator devtools0x08d6b0.RequestAccumulator ``
}

