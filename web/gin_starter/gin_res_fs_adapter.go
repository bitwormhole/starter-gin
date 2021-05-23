package gin_starter

import (
	"io/fs"

	"github.com/bitwormhole/starter/collection"
)

type GinFsAdapter struct {
	res collection.Resources
}

func (inst *GinFsAdapter) GetFS() fs.FS {
	return inst
}

func (inst *GinFsAdapter) Open(name string) (fs.File, error) {
	// todo ...
	return nil, nil
}
