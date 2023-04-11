package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "scm/db"
	"strconv"
)

func GetCourse(ctx *gin.Context) {
	var courses []Course
	DB.Preload("Teacher").Find(&courses)
	ctx.HTML(200, "course.html", gin.H{
		"courses": courses,
	})
}
func GetCourseHtml(ctx *gin.Context) {
	var teachers []Teacher
	DB.Find(&teachers)
	ctx.HTML(200, "addCourse.html", gin.H{
		"teachers": teachers,
	})
	fmt.Print(teachers[0].Name)
}
func AddCourse(ctx *gin.Context) {
	name := ctx.PostForm("name")
	credit, _ := strconv.Atoi(ctx.PostForm("credit"))
	period, _ := strconv.Atoi(ctx.PostForm("period"))
	teacher, _ := strconv.Atoi(ctx.PostForm("teacher"))
	addCourses := Course{Base: Base{Name: name}, Credit: credit, Period: period, TeacherID: teacher}
	DB.Create(&addCourses)
	ctx.Redirect(301, "/course")
}
