package route

import (
	"github.com/gin-gonic/gin"
	. "scm/core"
)

func InitTeachRoute(r *gin.RouterGroup) {
	r.GET("/", GetTeacher)
	//get请求拿页面，POST请求拿数据
	r.GET("/add", GetAddTeacherHtml)
	r.POST("/add", AddTeacher)
}
