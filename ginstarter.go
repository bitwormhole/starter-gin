package ginstarter

import (
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
)

// InitGin 函数：开始初始化Gin应用程序
func InitGin() application.Initializer {
	appinit := starter.InitApp()
	appinit.Use(Module())
	return appinit
}
