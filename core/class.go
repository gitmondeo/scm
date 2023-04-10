package core

import (
	"github.com/gin-gonic/gin"
	. "scm/db"
	"strconv"
)

func GetClass(ctx *gin.Context) {
	var classes []Class
	DB.Preload("Teacher").Find(&classes)
	ctx.HTML(200, "class.html", gin.H{
		"classes": classes,
	})
}
func GetClassHtml(ctx *gin.Context) {
	var tutors []Teacher
	DB.Find(&tutors)
	ctx.HTML(200, "addClass.html", gin.H{
		"tutors": tutors,
	})
}

func AddClass(ctx *gin.Context) {
	name := ctx.PostForm("name")
	num, _ := strconv.Atoi(ctx.PostForm("num"))
	tutor, _ := strconv.Atoi(ctx.PostForm("tutor"))
	classes := Class{Base: Base{Name: name}, Num: num, TutorID: tutor}
	DB.Create(&classes)
	//ctx.JSON(200, gin.H{
	//	"classes": classes,
	//})
	ctx.Redirect(301, "/class")

}
