package web

import (
	"github.com/bitwormhole/starter/application"
)

type GinServerContext struct {

	// public
	ApplicationContext application.Context
	Configuration      *GinServerConfig
	ContentTypeManager *ContentTypeManager
	ResourceLoader     *StaticResourceLoader

	// private
	runtime        *myGinRuntime
	indexPageNames []string
}
