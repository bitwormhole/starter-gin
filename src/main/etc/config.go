package etc

import "github.com/bitwormhole/starter/application"

// ExportConfig 导出 starter-gin 配置
func ExportGinConfig(cb application.ConfigBuilder) error {

	configProperties(cb)
	return autoGenConfig(cb)
}
