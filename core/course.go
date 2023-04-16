package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	. "scm/db"
	"strconv"
)

func GetCourse(ctx *gin.Context) {
	var courses []Course
	searchParams := ctx.Query("searchParams")
	if searchParams == "" {
		DB.Preload("Teacher").Find(&courses)
	} else {
		DB.Where("name like ?", "%"+searchParams+"%").Preload("Teacher").Find(&courses)
	}

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

//DeleteCourse 删除课程
func DeleteCourse(ctx *gin.Context) {
	delID := ctx.Param("delID")
	fmt.Println("delID:::", delID)
	DB.Where("name = ?", delID).Delete(&Course{})
	ctx.Redirect(http.StatusMovedPermanently, "/course")
}

//
func GetEditCourseHtml(ctx *gin.Context) {
	editID := ctx.Param("editID")
	var courses Course
	DB.Where("name = ?", editID).Find(&courses)
	//查出所有教师
	var teachers []Teacher
	DB.Find(&teachers)
	ctx.HTML(200, "editCourse.html", gin.H{
		"courses":  courses,
		"teachers": teachers,
	})

}

func EditCourse(ctx *gin.Context) {
	name := ctx.PostForm("name")
	credit, _ := strconv.Atoi(ctx.PostForm("credit"))
	period, _ := strconv.Atoi(ctx.PostForm("period"))
	teacher, _ := strconv.Atoi(ctx.PostForm("teacher"))
	DB.Model(&Course{}).Where("name = ?", name).Updates(&Course{Base: Base{}, Credit: credit, Period: period, TeacherID: teacher})
	ctx.Redirect(301, "/course")
}
