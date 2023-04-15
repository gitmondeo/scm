package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "scm/db"
	"strconv"
	//. "scm/db"
)

// GetClass 课程页面
func GetClass(ctx *gin.Context) {
	var classes []Class
	searchParams := ctx.Query("searchParams")
	if searchParams == "" {
		DB.Find(&classes)
	} else {
		DB.Where("name like ?", "%"+searchParams+"%").Find(&classes)
	}
	ctx.HTML(200, "class.html", gin.H{
		"classes": classes,
	})
}

//添加班级页面
func GetClassHtml(ctx *gin.Context) {
	var tutors []Teacher
	DB.Find(&tutors)
	ctx.HTML(200, "addClass.html", gin.H{
		"tutors": tutors,
	})
}

//添加班级
func AddClass(ctx *gin.Context) {
	name := ctx.PostForm("name")
	num, _ := strconv.Atoi(ctx.PostForm("num"))
	tutor, _ := strconv.Atoi(ctx.PostForm("tutor"))
	fmt.Println(tutor)
	addClass := Class{Base: Base{Name: name}, Num: num, TutorID: tutor}
	fmt.Println(addClass)
	DB.Create(&addClass)
	ctx.Redirect(301, "/class")
}
