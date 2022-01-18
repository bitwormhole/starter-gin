package config_lib

import application "github.com/bitwormhole/starter/application"

func ExportConfigForGinStarterLib(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
