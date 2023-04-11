package core

import (
	"github.com/gin-gonic/gin"
	. "scm/db"
	"strconv"
)

func GetTeacher(ctx *gin.Context) {
	var teachers []Teacher
	DB.Preload("Class").Find(&teachers)
	ctx.HTML(200, "teacher.html", gin.H{
		"teachers": teachers,
	})
}
func GetAddTeacherHtml(ctx *gin.Context) {
	var classes []Class
	DB.Find(&classes)
	ctx.HTML(200, "addTeacher.html", gin.H{
		"classes": classes,
	})
}
func AddTeacher(ctx *gin.Context) {
	//获取前端请求数据
	sno, _ := strconv.Atoi(ctx.PostForm("sno"))
	name := ctx.PostForm("name")
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	gender := ctx.PostForm("gender")
	tel := ctx.PostForm("tel")
	cls, _ := strconv.Atoi(ctx.PostForm("cls"))
	remark := ctx.PostForm("remark")
	//赋值给student对象
	stus := Student{Base: Base{Name: name}, Sno: sno, Age: age, Gender: gender, Tel: tel, ClassID: cls, Remark: remark}
	//数据库存储
	DB.Create(&stus)
	//数据库查询
	ctx.Redirect(301, "/teacher")

}
