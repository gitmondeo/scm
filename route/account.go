package route

import (
	"github.com/gin-gonic/gin"
	. "scm/core"
)

func InitAccountRoute(r *gin.RouterGroup) {
	r.GET("/", GetAccount)
}
