package starter

import (
	"github.com/bitwormhole/starter-gin/etc/glass"
	s2 "github.com/bitwormhole/starter/starter"

	"github.com/bitwormhole/starter/application"
)

// Gin 函数：开始初始化Gin应用程序
func Gin() application.Initializer {
	appinit := s2.InitApp()
	appinit.Use(glass.ExportModule())
	return appinit
}
