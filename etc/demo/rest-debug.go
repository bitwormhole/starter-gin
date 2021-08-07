package demo

import (
	"github.com/bitwormhole/starter-gin/glass"
	application "github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

type theDebugController struct {
	markup.Component
	instance *glass.DebugInfoController `class:"rest-controller"`

	AppContext application.Context `inject:"context"`
}
