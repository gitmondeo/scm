package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "scm/db"
	"strconv"
)

// GetTeacher 教师页面
func GetTeacher(ctx *gin.Context) {
	var teachers []Teacher
	searchParams := ctx.Query("searchParams")
	if searchParams == "" {
		DB.Find(&teachers)
	} else {
		DB.Where("name like ?", "%"+searchParams+"%").Find(&teachers)
	}

	ctx.HTML(200, "teacher.html", gin.H{
		"teachers": teachers,
	})
}

// GetAddTeacherHtml 获取添加教师界面
func GetAddTeacherHtml(ctx *gin.Context) {
	ctx.HTML(200, "addTeacher.html", nil)
}

// AddTeacher 添加教师
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

// DeleteTeacher 删除教师
func DeleteTeacher(ctx *gin.Context) {
	delID := ctx.Param("delID")
	fmt.Println("delID:::", delID)
	DB.Where("tno = ?", delID).Delete(&Teacher{})
	ctx.Redirect(301, "/teacher")
}
