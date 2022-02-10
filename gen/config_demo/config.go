package config_demo

import "github.com/bitwormhole/starter/application"

func ExportConfigForStarterGinDemo(cb application.ConfigBuilder) error {

	return autoGenConfig(cb)
}
