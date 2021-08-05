package srctestgo

import (
	"testing"

	root "github.com/bitwormhole/starter-gin"
	"github.com/bitwormhole/starter-gin/starter"
)

func TestGinStarter(t *testing.T) {
	starter.Gin().MountResources(root.GetResources(), "/").Run()
}
