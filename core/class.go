package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "scm/db"
	"strconv"
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
	var tutors []Teacher
	DB.Find(&tutors)
	ctx.HTML(200, "class.html", gin.H{
		"classes": classes,
		"tutors":  tutors,
	})
}

// GetClassHtml 添加班级页面
func GetClassHtml(ctx *gin.Context) {
	var tutors []Teacher
	DB.Find(&tutors)
	ctx.HTML(200, "addClass.html", gin.H{
		"tutors": tutors,
	})
}

// AddClass 添加班级
func AddClass(ctx *gin.Context) {
	name := ctx.PostForm("name")
	num, _ := strconv.Atoi(ctx.PostForm("num"))
	tutor, _ := strconv.Atoi(ctx.PostForm("tutor"))
	addClass := Class{Base: Base{Name: name}, Num: num, TutorID: tutor}
	DB.Create(&addClass)
	ctx.Redirect(301, "/class")
}

// DeleteClass 删除班级
func DeleteClass(ctx *gin.Context) {
	delID := ctx.Param("delID")
	DB.Where("name = ?", delID).Delete(&Class{})
	ctx.Redirect(http.StatusMovedPermanently, "/class")
}

// GetEditClassHtml 获取编辑班级页面
func GetEditClassHtml(ctx *gin.Context) {
	var tutors []Teacher
	DB.Find(&tutors)

	editID := ctx.Param("editID")
	var classes Class
	DB.Where("name = ?", editID).Find(&classes)
	ctx.HTML(200, "editClass.html", gin.H{
		"classes": classes,
		"tutors":  tutors,
	})
}

// EditClass 编辑班级
func EditClass(ctx *gin.Context) {
	name := ctx.PostForm("name")
	num, _ := strconv.Atoi(ctx.PostForm("num"))
	tutor, _ := strconv.Atoi(ctx.PostForm("tutor"))
	DB.Model(&Class{}).Where("name = ?", name).Updates(&Class{Base: Base{}, Num: num, TutorID: tutor})
	ctx.Redirect(301, "/class")
}
