package render

import "github.com/gin-contrib/multitemplate"

// CreateMyRender 加载第三方模块文件
func CreateMyRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	render.AddFromFiles("base.html", "templates/base/base.html")
	render.AddFromFiles("student.html", "templates/base/base.html", "templates/student/student.html")
	render.AddFromFiles("class.html", "templates/base/base.html", "templates/class/class.html")
	render.AddFromFiles("course.html", "templates/base/base.html", "templates/course/course.html")
	render.AddFromFiles("teacher.html", "templates/base/base.html", "templates/teacher/teacher.html")

	//student
	render.AddFromFiles("addStudent.html", "templates/base/base.html", "templates/student/addStudent.html")
	render.AddFromFiles("editStudent.html", "templates/base/base.html", "templates/student/editStudent.html")
	render.AddFromFiles("detailStudent.html", "templates/base/base.html", "templates/student/detailStudent.html")
	render.AddFromFiles("selectCourse.html", "templates/base/base.html", "templates/student/selectCourse.html")
	render.AddFromFiles("editDetailStudent.html", "templates/base/base.html", "templates/student/editDetailStudent.html")

	//class
	render.AddFromFiles("addClass.html", "templates/base/base.html", "templates/class/addClass.html")
	render.AddFromFiles("editClass.html", "templates/base/base.html", "templates/class/editClass.html")
	//course
	render.AddFromFiles("addCourse.html", "templates/base/base.html", "templates/course/addCourse.html")
	render.AddFromFiles("editCourse.html", "templates/base/base.html", "templates/course/editCourse.html")
	//teacher
	render.AddFromFiles("addTeacher.html", "templates/base/base.html", "templates/teacher/addTeacher.html")
	render.AddFromFiles("editTeacher.html", "templates/base/base.html", "templates/teacher/editTeacher.html")
	//login
	render.AddFromFiles("login.html", "templates/login/login.html")
	//account登录账号
	render.AddFromFiles("account.html", "templates/base/base.html", "templates/account/account.html")
	render.AddFromFiles("addAccount.html", "templates/base/base.html", "templates/account/addAccount.html")
	render.AddFromFiles("editAccount.html", "templates/base/base.html", "templates/account/editAccount.html")
	//返回render
	return render

}
