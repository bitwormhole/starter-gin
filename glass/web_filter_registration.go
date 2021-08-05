package glass

import "github.com/gin-gonic/gin"

type innerFilterRegistration struct {
	order int
	fn    gin.HandlerFunc
}
