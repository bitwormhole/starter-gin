package config_devtools

import "github.com/bitwormhole/starter/application"

func ExportConfigForGinStarterDevtools(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
