package core

import (
	"github.com/gin-gonic/gin"
	. "scm/db"
	"strconv"
)

func GetTeacher(ctx *gin.Context) {
	var teachers []Teacher
	DB.Find(&teachers)
	ctx.HTML(200, "teacher.html", gin.H{
		"teachers": teachers,
	})
}
func GetAddTeacherHtml(ctx *gin.Context) {
	ctx.HTML(200, "addTeacher.html", nil)
}
func AddTeacher(ctx *gin.Context) {
	//获取前端请求数据
	tno, _ := strconv.Atoi(ctx.PostForm("tno"))
	name := ctx.PostForm("name")
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	gender := ctx.PostForm("gender")
	tel := ctx.PostForm("tel")
	remark := ctx.PostForm("remark")
	//赋值给student对象
	teachers := Teacher{Base: Base{Name: name}, Tno: tno, Age: age, Gender: gender, Tel: tel, Remark: remark}
	//数据库存储
	DB.Create(&teachers)
	//数据库查询
	ctx.Redirect(301, "/teacher")

}
