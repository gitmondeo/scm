package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	. "scm/db"
	"strconv"
)

// GetALLStudent 学生页面，模糊搜索
func GetALLStudent(ctx *gin.Context) {
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

// GetAddStuHtml 添加学生页面
func GetAddStuHtml(ctx *gin.Context) {
	var classes []Class
	DB.Find(&classes)

	var users []UserInfo
	DB.Find(&users)

	fmt.Println("userinfo:::", users)
	ctx.HTML(200, "addStudent.html", gin.H{
		"classes": classes,
		"users":   users,
	})
}

// AddStudent 添加学生，保存到数据库
func AddStudent(ctx *gin.Context) {
	//获取前端请求数据
	sno, _ := strconv.Atoi(ctx.PostForm("sno"))
	name := ctx.PostForm("name")
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	gender := ctx.PostForm("gender")
	tel := ctx.PostForm("tel")
	cls, _ := strconv.Atoi(ctx.PostForm("cls"))
	remark := ctx.PostForm("remark")
	//pwd := ctx.PostForm("pwd")
	userinfoID, _ := strconv.Atoi(ctx.PostForm("userinfoID"))
	//赋值给student对象
	stus := Student{Base: Base{Name: name}, Sno: sno, Age: age, Gender: gender, Tel: tel, ClassID: cls, UserInfoID: userinfoID, Remark: remark}
	//数据库存储
	DB.Create(&stus)
	//添加一个学生，班级人数加1，通过gorm.Expr表达式实现
	DB.Model(&Class{}).Where("id = ?", cls).Update("num", gorm.Expr("num+1"))
	//数据库查询
	ctx.Redirect(301, "/student")

}

// DeleteStudent 删除学生
func DeleteStudent(ctx *gin.Context) {
	delID := ctx.Param("delID")
	fmt.Println("delID", delID)
	DB.Where("sno = ?", delID).Delete(&Student{})
	//ctx.String(200, "ok")
	ctx.Redirect(http.StatusMovedPermanently, "/student")
}

// GetEditStuHtml 编辑学生页面
func GetEditStuHtml(ctx *gin.Context) {
	var students Student
	editID := ctx.Param("editID")
	DB.Where("sno = ?", editID).Find(&students)

	//查找出所有班级
	var classes []Class
	DB.Find(&classes)
	ctx.HTML(200, "editStudent.html", gin.H{
		"students": students,
		"classes":  classes,
	})
}

// EditStudent 编辑学生
func EditStudent(ctx *gin.Context) {
	//获取前端请求数据
	sno, _ := strconv.Atoi(ctx.PostForm("sno"))
	name := ctx.PostForm("name")
	age, _ := strconv.Atoi(ctx.PostForm("age"))
	gender := ctx.PostForm("gender")
	tel := ctx.PostForm("tel")
	cls, _ := strconv.Atoi(ctx.PostForm("cls"))
	remark := ctx.PostForm("remark")
	pwd := ctx.PostForm("pwd")

	//查询sno对应的ID
	//1)查找出所有学生
	var students []Student
	DB.Find(&students)
	//2）根据sno找到id
	var ID int
	for _, stu := range students {
		if stu.Sno == sno {
			ID = stu.ID
		}
	}
	//修改数据库，Model是条件，默认是ID主键值才行，下方是自定义条件
	DB.Model(&Student{}).Where("ID = ?", ID).Updates(Student{Base: Base{Name: name}, Sno: sno, Age: age, Gender: gender, Tel: tel, ClassID: cls, Remark: remark, Pwd: pwd})
	ctx.Redirect(http.StatusMovedPermanently, "/student")
}

// GetOneStudent 获取单个学生页面
func GetOneStudent(ctx *gin.Context) {
	sno := ctx.Param("sno")
	var student Student
	DB.Preload("Class").Where("sno = ?", sno).Find(&student)

	//查询学生选择的课程
	var selectCourses []Course
	DB.Preload("Teacher").Model(&student).Association("Course").Find(&selectCourses)
	ctx.HTML(http.StatusOK, "detailStudent.html", gin.H{
		"student":       student,
		"selectCourses": selectCourses,
	})
}

// GetSelectCourseHtml 获取学生选课界面
func GetSelectCourseHtml(ctx *gin.Context) {
	sno := ctx.Param("sno")
	var student Student
	DB.Where("sno = ?", sno).Find(&student)
	var courses []Course
	DB.Preload("Teacher").Find(&courses)
	ctx.HTML(http.StatusOK, "selectCourse.html", gin.H{
		"courses": courses,
		"student": student,
	})
}

func SelectCourse(ctx *gin.Context) {
	sno := ctx.Param("sno")
	var student Student
	DB.Where("sno = ?", sno).Find(&student)

	courseID := ctx.PostFormArray("course_id")
	var courses []Course
	DB.Where("id in ?", courseID).Find(&courses)

	DB.Preload("Teacher").Model(&student).Association("Course").Append(&courses)
	ctx.Redirect(http.StatusMovedPermanently, "/student/"+sno)
}

func DeleteSelectCourse(ctx *gin.Context) {
	//根据学号，获取该学生对象,得到学生的ID
	sno := ctx.Param("sno")
	var student Student
	DB.Where("sno = ?", sno).Find(&student)
	student_id := student.ID
	//获取要删除的已选课程的ID
	courseID := ctx.Param("courseID")
	//取消学生和课程的绑定
	DB.Where("student_id = ? AND course_id = ?", student_id, courseID).Model(&student).Association("Course").Clear()
	//DB.Model(&student).Association("Course").Clear()
	ctx.Redirect(http.StatusMovedPermanently, "/student/"+sno)
}
