package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	var users []UserInfo
	DB.Find(&users)
	ctx.HTML(200, "addTeacher.html", gin.H{
		"users": users,
	})
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
	userinfoID, _ := strconv.Atoi(ctx.PostForm("userinfoID"))
	//pwd := ctx.PostForm("pwd")
	//赋值给student对象
	teachers := Teacher{Base: Base{Name: name}, Tno: tno, Age: age, Gender: gender, Tel: tel, UserInfoID: userinfoID, Remark: remark}
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
	ctx.Redirect(http.StatusMovedPermanently, "/teacher")
}

// GetEditTeacherHtml 获取编辑教师页面
func GetEditTeacherHtml(ctx *gin.Context) {
	editID := ctx.Param("editID")
	var teachers Teacher
	DB.Where("tno = ?", editID).Find(&teachers)
	ctx.HTML(200, "editTeacher.html", gin.H{
		"teachers": teachers,
	})
}

//EditTeacher 编辑教师
func EditTeacher(ctx *gin.Context) {
	//获取前端请求数据
	tno, _ := strconv.Atoi(ctx.PostForm("tno"))
	name := ctx.PostForm("name")
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	gender := ctx.PostForm("gender")
	tel := ctx.PostForm("tel")
	remark := ctx.PostForm("remark")
	//pwd := ctx.PostForm("pwd")
	//根据tno查找id
	//fmt.Println("params:::", tel, tno, name, age, gender, remark, pwd)
	var teacher Teacher
	DB.Where("tno = ?", tno).Find(&teacher)
	DB.Model(&Teacher{}).Where("id = ?", teacher.ID).Updates(&Teacher{Base: Base{Name: name}, Gender: gender, Age: age, Tel: tel, Remark: remark})
	ctx.Redirect(http.StatusMovedPermanently, "/teacher")
}
