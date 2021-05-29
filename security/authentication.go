package security

import "github.com/bitwormhole/starter/lang"

// Authentication (认证) token
type AuthenticationToken interface {

	// 委托人
	Principal() lang.Object

	// 凭证
	Credentials() lang.Object
}

// Authentication (认证) info
type AuthenticationInfo interface {

	// 委托人
	Principal() lang.Object

	Roles() []string
	Permissions() []string
	Success() bool
	Message() string
}

// Authentication 服务
type AuthenticationService interface {
	Name() string
	Supports(token AuthenticationToken) bool
	Authenticate(token AuthenticationToken) (AuthenticationInfo, error)
}
