package glass

import (
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

type ContentTypeManager interface {
	Find(pathname string) string
}

////////////////////////////////////////////////////////////////////////////////

type DefaultContentTypeManager struct {
	AppContext      application.Context
	TypesProperties string

	// private
	table collection.Properties
}

func (inst *DefaultContentTypeManager) _Impl() ContentTypeManager {
	return inst
}

func (inst *DefaultContentTypeManager) Find(name string) string {

	// TODO : ...
	return "todo..."
}

////////////////////////////////////////////////////////////////////////////////
