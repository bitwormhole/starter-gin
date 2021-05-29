package security

// Authorization 授权
type Authorization interface {
	Roles() []string
	Permissions() []string
}

// Authorization 服务
type AuthorizationService interface {
}
