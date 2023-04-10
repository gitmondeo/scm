package core

import (
	"github.com/gin-gonic/gin"
	. "scm/db"
	"strconv"
)

func GetCourse(ctx *gin.Context) {
	ctx.HTML(200, "course.html", nil)
}
func GetCourseHtml(ctx *gin.Context) {
	ctx.HTML(200, "addCourse.html", nil)
}
func AddCourse(ctx *gin.Context) {
	name := ctx.PostForm("name")
	credit, _ := strconv.Atoi(ctx.PostForm("credit"))
	period, _ := strconv.Atoi(ctx.PostForm("period"))
	teacher, _ := strconv.Atoi(ctx.PostForm("teacher"))
	courses := Course{Base: Base{Name: name}, Credit: credit, Period: period, TeacherID: teacher}
	DB.Create(&courses)
	ctx.JSON(200, gin.H{
		"courses": courses,
	})
}
