package devtools

import (
	srcdevtools "github.com/bitwormhole/starter-gin/src/devtools"
	application "github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// ExportConfig 导出配置
func ExportConfig(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}

type theWebDevtoolsController struct {
	markup.Component
	instance *srcdevtools.DevtoolsController `class:"rest-controller"`

	AppContext application.Context `inject:"context"`
	Enable     bool                `inject:"${web.devtools.enable}"`
}
