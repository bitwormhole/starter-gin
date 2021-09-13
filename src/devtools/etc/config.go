package etc

import "github.com/bitwormhole/starter/application"

// ExportDevtoolsConfig 导出 devtools 的配置
func ExportDevtoolsConfig(cb application.ConfigBuilder) error {

	return autoGenConfig(cb)
}
