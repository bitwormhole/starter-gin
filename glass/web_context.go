package glass

// WebContext 表示一个web上下文(service)
type WebContext interface {
	GetContextID() string
	GetContainer() Container
	GetContextPath() string
	GetControllers() []Controller
	GetInterceptors() []InterceptorRegistry
}
