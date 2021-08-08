package startergin

import (
	"github.com/bitwormhole/starter-gin/etc/glass"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/starter"
)

// InitGinApp 函数：开始初始化Gin应用程序
func InitGinApp() application.Initializer {
	appinit := starter.InitApp()
	appinit.Use(Module())
	return appinit
}

// Module 函数：返回 starter-gin (release)模块
func Module() application.Module {
	return glass.ExportModuleRelease()
}

// ModuleDebug 函数：返回 starter-gin (debug) 模块
func ModuleDebug() application.Module {
	return glass.ExportModuleDebug()
}
