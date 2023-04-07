package route

import (
	"github.com/gin-gonic/gin"
	. "scm/core"
)

func InitClassRoute(r *gin.RouterGroup) {
	r.GET("/", GetClass)
	r.GET("/add", GetClassHtml)
}
