package route

import (
	"github.com/gin-gonic/gin"
	. "scm/core"
)

func InitCourseRoute(r *gin.RouterGroup) {
	r.GET("/", GetCourse)
	r.GET("/add", GetCourseHtml)
	r.POST("/add", AddCourse)
}
