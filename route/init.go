package route

import (
	"github.com/gin-gonic/gin"
	. "scm/core"
)

func InitRoute(r *gin.Engine) {
	//登录请求
	r.GET("/login", GetLoginHtml)
	r.POST("/login", Login)

	//首页请求
	r.GET("/", GetIndex)

	//student请求
	student := r.Group("/student")
	InitStuRoute(student)

	//class请求
	class := r.Group("/class")
	InitClassRoute(class)

	//课程请求
	course := r.Group("/course")
	InitCourseRoute(course)

	//teacher请求
	teacher := r.Group("/teacher")
	InitTeachRoute(teacher)

	//account请求
	account := r.Group("/account")
	InitAccountRoute(account)

}
