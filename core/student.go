package core

import (
	"github.com/gin-gonic/gin"
	. "scm/db"
	"strconv"
)

func GetStudent(ctx *gin.Context) {
	ctx.HTML(200, "student.html", nil)
}
func GetAddStuHtml(ctx *gin.Context) {
	ctx.HTML(200, "addStudent.html", nil)
}
func AddStudent(ctx *gin.Context) {
	//获取前端请求数据
	sno, _ := strconv.Atoi(ctx.PostForm("sno"))
	name := ctx.PostForm("name")
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	gender := ctx.PostForm("gender")
	tel := ctx.PostForm("tel")
	cls, _ := strconv.Atoi(ctx.PostForm("cls"))
	//赋值给student对象
	stus := Student{Base: Base{ID: 1, Name: name}, Sno: sno, Age: age, Gender: gender, Tel: tel, ClassID: cls}
	//数据库存储
	DB.Create(&stus)
	//数据库查询
	ctx.String(200, "添加完成")
}
