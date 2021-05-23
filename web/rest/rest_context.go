package rest

const ContextClassName = "rest-context"

type Context interface {
	ForType(typeName string) Type
}

type Type interface {
	Path(path string) string
}
