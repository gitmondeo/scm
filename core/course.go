package core

import "github.com/gin-gonic/gin"

func GetCourse(ctx *gin.Context) {
	ctx.HTML(200, "course.html", nil)
}
func GetCourseHtml(ctx *gin.Context) {
	ctx.HTML(200, "addCourse.html", nil)
}
func AddCourse(ctx *gin.Context) {
	ctx.HTML(200, "addCourse.html", nil)
}
