package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	. "scm/db"
	"strconv"
)

//学生页面，模糊搜索
func GetStudent(ctx *gin.Context) {
	var students []Student
	//获取搜索框传入的参数
	searchParams := ctx.Query("searchParams")
	//如果搜索框有参数，进行模糊搜索，否则搜索所有内容（%searchParams% 两个百分号进行模糊搜索）
	if searchParams == "" {
		DB.Preload("Class").Find(&students)
	} else {
		DB.Where("name like ?", "%"+searchParams+"%").Preload("Class").Find(&students)
	}

	ctx.HTML(200, "student.html", gin.H{
		"students": students,
	})
}

//添加学生页面
func GetAddStuHtml(ctx *gin.Context) {
	var classes []Class
	DB.Find(&classes)
	ctx.HTML(200, "addStudent.html", gin.H{
		"classes": classes,
	})
}

//添加学生，保存到数据库
func AddStudent(ctx *gin.Context) {
	//获取前端请求数据
	sno, _ := strconv.Atoi(ctx.PostForm("sno"))
	name := ctx.PostForm("name")
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	gender := ctx.PostForm("gender")
	tel := ctx.PostForm("tel")
	cls, _ := strconv.Atoi(ctx.PostForm("cls"))
	remark := ctx.PostForm("remark")
	pwd := ctx.PostForm("pwd")
	//赋值给student对象
	stus := Student{Base: Base{Name: name}, Sno: sno, Age: age, Gender: gender, Tel: tel, Pwd: pwd, ClassID: cls, Remark: remark}
	//数据库存储
	DB.Create(&stus)
	//数据库查询
	ctx.Redirect(301, "/student")

}

//删除学生
func DeleteStudent(ctx *gin.Context) {
	delID := ctx.Param("delID")
	fmt.Println("delID", delID)
	DB.Where("sno = ?", delID).Delete(&Student{})
	//ctx.String(200, "ok")
	ctx.Redirect(http.StatusMovedPermanently, "/student")
}

//编辑学生页面
func GetEditStuHtml(ctx *gin.Context) {
	var students Student
	editID := ctx.Param("editID")
	DB.Where("sno = ?", editID).Preload("Class").Find(&students)
	ctx.HTML(200, "editStudent.html", gin.H{
		"students": students,
	})
}

//编辑学生
func EditStudent(ctx *gin.Context) {

	ctx.Redirect(http.StatusMovedPermanently, "/student")
}
