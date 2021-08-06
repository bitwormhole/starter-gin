package initd

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
)

// FileSystemResources 从本地文件系统加载资源
func FileSystemResources(path string) collection.Resources {
	inst := &filesysResources{}
	return inst.init(path)
}

////////////////////////////////////////////////////////////////////////////////

type filesysResources struct {
	base fs.Path
}

func (inst *filesysResources) _Impl() collection.Resources {
	return inst
}

func (inst *filesysResources) init(path string) collection.Resources {

	const prefixEnv = "env:"
	const prefixArgs = "args:"
	const prefixFile = "file:"

	if strings.HasPrefix(path, prefixFile) {
		path = path[len(prefixFile):]
	} else if strings.HasPrefix(path, prefixArgs) {
		key := path[len(prefixArgs):]
		path = inst.getBasePathFromArgs(key)
	} else if strings.HasPrefix(path, prefixEnv) {
		key := path[len(prefixEnv):]
		path = inst.getBasePathFromEnv(key)
	} else {
		// NOP
	}

	inst.base = fs.Default().GetPath(path)
	return inst
}

func (inst *filesysResources) getBasePathFromArgs(key string) string {
	key = "--" + key
	args := collection.CreateArguments()
	args.Import(os.Args)
	rd, ok1 := args.GetReader(key)
	if !ok1 {
		panic("no argument:" + key)
	}
	for {
		value, ok := rd.Read()
		if !ok {
			break
		}
		if value == key {
			continue
		}
		return value
	}
	panic("no argument:" + key)
}

func (inst *filesysResources) getBasePathFromEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("no env:" + key)
	}
	return value
}

func (inst *filesysResources) mapping(resPath string) fs.Path {
	const token = ":/"
	index := strings.Index(resPath, token)
	if index > 0 {
		resPath = resPath[index+len(token):]
	}
	return inst.base.GetChild("./" + resPath)
}

func (inst *filesysResources) GetText(path string) (string, error) {
	file := inst.mapping(path)
	return file.GetIO().ReadText()
}

func (inst *filesysResources) GetBinary(path string) ([]byte, error) {
	file := inst.mapping(path)
	return file.GetIO().ReadBinary()
}

func (inst *filesysResources) GetReader(path string) (io.ReadCloser, error) {
	// file := inst.mapping(path)
	// return file.GetIO()
	return nil, errors.New("no impl")
}

func (inst *filesysResources) All() []*collection.Resource {
	return inst.List("/", false)
}

func (inst *filesysResources) List(path string, recursive bool) []*collection.Resource {
	walker := &filesysResWalker{}
	walker.owner = inst
	walker.recursive = recursive
	walker.basepath = path
	return walker.walk()
}

////////////////////////////////////////////////////////////////////////////////

type filesysResWalker struct {
	owner     *filesysResources
	recursive bool
	basepath  string

	results []*collection.Resource
}

func (inst *filesysResWalker) walk() []*collection.Resource {
	inst.results = make([]*collection.Resource, 0)
	root := inst.root()
	inst.innerWalk(root, 99, false)
	return inst.results
}

func (inst *filesysResWalker) child(name string, parent *collection.Resource) *collection.Resource {
	ch := &collection.Resource{}
	ch.Name = name
	ch.IsDir = false
	if parent == nil {
		// root
		ch.BasePath = inst.basepath
		ch.AbsolutePath = inst.basepath
		ch.RelativePath = ""
	} else {
		// child
		ch.BasePath = parent.BasePath
		ch.AbsolutePath = parent.AbsolutePath + "/" + name
		ch.RelativePath = parent.RelativePath + "/" + name
	}
	return ch
}

func (inst *filesysResWalker) root() *collection.Resource {
	return inst.child("", nil)
}

func (inst *filesysResWalker) innerWalk(node *collection.Resource, limit int, deep bool) error {
	if limit < 1 {
		return errors.New("path is too deep, path=" + node.AbsolutePath)
	}

	path := inst.owner.mapping(node.AbsolutePath)
	node.IsDir = path.IsDir()

	if path.IsDir() {
		inst.onDir(node)
	} else if path.IsFile() {
		inst.onFile(node)
		return nil
	} else {
		return nil
	}
	if deep && !inst.recursive {
		return nil
	}
	names := path.ListNames()
	for _, name := range names {
		child := inst.child(name, node)
		err := inst.innerWalk(child, limit-1, true)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *filesysResWalker) onDir(node *collection.Resource) {
	inst.results = append(inst.results, node)
}

func (inst *filesysResWalker) onFile(node *collection.Resource) {
	inst.results = append(inst.results, node)
}

////////////////////////////////////////////////////////////////////////////////
