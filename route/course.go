package route

import (
	"github.com/gin-gonic/gin"
	. "scm/core"
)

func InitCourseRoute(r *gin.RouterGroup) {
	r.GET("/", GetCourse)
}
