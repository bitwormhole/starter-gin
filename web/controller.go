package web

type Controller interface {
	Config(c Container) error
}
