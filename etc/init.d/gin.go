package initd

import (
	"github.com/bitwormhole/starter-gin/etc/glass"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/starter"
)

// GinStarter 获取 gin-starter 启动器
func GinStarter() application.Initializer {
	ai := starter.InitApp()
	ai.Use(GinModule())
	return ai
}

// GinModule 获取 gin-starter 模块
func GinModule() application.Module {
	return glass.ExportModule()
}
