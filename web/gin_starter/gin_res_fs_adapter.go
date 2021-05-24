package gin_starter

import (
	"errors"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/bitwormhole/starter/collection"
)

type GinFsAdapter struct {
	res collection.Resources
}

func (inst *GinFsAdapter) GetFS() http.FileSystem {
	return inst
}

func (inst *GinFsAdapter) Open(name string) (http.File, error) {
	// todo ...

	path := inst.getPathByName(name)
	data, err := inst.res.GetBinary(path)
	if err != nil {
		return nil, err
	}

	file := &ginFsAdapterFile{}
	err = file.init(name, data)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (inst *GinFsAdapter) getPathByName(name string) string {

	if strings.HasSuffix(name, "/") {
		name = name + "index.html"
	}

	if strings.HasPrefix(name, "/") {
		return "www" + name
	} else {
		return "www/" + name
	}
}

////////////////////////////////////////////////////////////////////////////////

type ginFsAdapterFile struct {
	name   string
	data   []byte
	reader int
	size   int64
	info   fs.FileInfo
}

func (inst *ginFsAdapterFile) init(name string, data []byte) error {
	inst.name = name
	inst.data = data
	inst.reader = 0
	inst.size = int64(len(data))
	inst.info = &ginFsAdapterFileInfo{file: inst}
	return nil
}

func (inst *ginFsAdapterFile) _impl_all_() http.File {
	return inst
}

func (inst *ginFsAdapterFile) Stat() (fs.FileInfo, error) {
	return inst.info, nil
}

func (inst *ginFsAdapterFile) Seek(int64, int) (int64, error) {
	return 0, nil
}

func (inst *ginFsAdapterFile) Readdir(count int) ([]fs.FileInfo, error) {
	return nil, nil
}

func (inst *ginFsAdapterFile) Read(dst []byte) (int, error) {

	src := inst.data
	if src == nil || dst == nil {
		return 0, errors.New("file is closed.")
	}

	reader := inst.reader
	readerEnding := int(inst.size)
	writer := 0
	writerEnding := len(dst)
	count := 0

	for (reader < readerEnding) && (writer < writerEnding) {
		dst[writer] = src[reader]
		count++
		reader++
		writer++
	}

	return count, nil
}

func (inst *ginFsAdapterFile) Close() error {
	//	inst.data = nil
	inst.reader = 0
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type ginFsAdapterFileInfo struct {
	file *ginFsAdapterFile
}

func (inst *ginFsAdapterFileInfo) Name() string {
	return inst.file.name
}

func (inst *ginFsAdapterFileInfo) Size() int64 {
	// length in bytes for regular files; system-dependent for others
	return inst.file.size
}

func (inst *ginFsAdapterFileInfo) Mode() fs.FileMode {
	// file mode bits
	return 0
}

func (inst *ginFsAdapterFileInfo) ModTime() time.Time {
	// modification time
	return time.Unix(0, 0)
}

func (inst *ginFsAdapterFileInfo) IsDir() bool {
	// abbreviation for Mode().IsDir()
	return false
}

func (inst *ginFsAdapterFileInfo) Sys() interface{} {
	// underlying data source (can return nil)
	return nil
}
