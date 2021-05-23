package rest

import "strings"

////////////////////////////////////////////////////////////////////////////////

type SimpleRestContext struct {
	ContextPath string // like '/api/v1'
}

func (inst *SimpleRestContext) _impl_rest_context() Context {
	return inst
}

func (inst *SimpleRestContext) ForType(typeName string) Type {
	path := inst.ContextPath + "/" + typeName
	rt := &simpleRestType{path: path}
	return rt
}

////////////////////////////////////////////////////////////////////////////////

type simpleRestType struct {
	path string
}

func (inst *simpleRestType) _impl_rest_type() Type {
	return inst
}

func (inst *simpleRestType) Path(path string) string {
	if path == "" {
		return inst.path
	} else if strings.HasPrefix(path, "/") {
		return path
	}
	return inst.path + "/" + path
}
