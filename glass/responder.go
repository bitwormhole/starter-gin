package glass

import (
	"github.com/gin-gonic/gin"
)

// Response ...
type Response struct {
	Context *gin.Context
	Data    interface{}
	Status  int
	Error   error
}

// Responder ...
type Responder interface {
	Send(resp *Response)
}

// MainResponder ...
// 【inject:"#glass-main-responder"】
type MainResponder interface {
	Responder
	TargetResponderName() string
}

// ResponderRegistration ...
type ResponderRegistration struct {
	Name      string
	Responder Responder
}

// ResponderRegistry ...
// 【inject:".glass-responder-registry"】
type ResponderRegistry interface {
	GetRegistrationList() []*ResponderRegistration
}
