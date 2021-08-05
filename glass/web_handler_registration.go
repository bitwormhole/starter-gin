package glass

import "github.com/gin-gonic/gin"

type innerHandlerRegistration struct {
	method string
	path   string
	fn     gin.HandlerFunc
}
