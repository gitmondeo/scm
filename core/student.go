package core

import (
	"github.com/gin-gonic/gin"
	. "scm/db"
	"strconv"
)

func GetStudent(ctx *gin.Context) {
	var students []Student

	//DB.Find(&students)
	DB.Preload("Class").Find(&students)
	ctx.HTML(200, "student.html", gin.H{
		"students": students,
	})
}
func GetAddStuHtml(ctx *gin.Context) {
	var classes []Class
	DB.Find(&classes)
	ctx.HTML(200, "addStudent.html", gin.H{
		"classes": classes,
	})
}
func AddStudent(ctx *gin.Context) {
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
	var students []Student
	//DB.Find(&students)
	DB.Preload("Class").Find(&students)
	ctx.JSON(200, gin.H{
		"students": students,
	})
}
