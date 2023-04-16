package route

import (
	"github.com/gin-gonic/gin"
	. "scm/core"
)

func InitCourseRoute(r *gin.RouterGroup) {
	r.GET("/", GetCourse)
	r.GET("/add", GetCourseHtml)
	r.POST("/add", AddCourse)
	r.GET("/delete/:delID", DeleteCourse)
	r.GET("/edit/:editID", GetEditCourseHtml)
	r.POST("/edit/:editID", EditCourse)

}
