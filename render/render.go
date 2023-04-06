package render

import "github.com/gin-contrib/multitemplate"

// CreateMyRender 加载第三方模块文件
func CreateMyRender() multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	render.AddFromFiles("base.html", "templates/base/base.html")
	render.AddFromFiles("student.html", "templates/base/base.html", "templates/student/student.html")
	render.AddFromFiles("class.html", "templates/base/base.html", "templates/class/class.html")
	render.AddFromFiles("course.html", "templates/base/base.html", "templates/course/course.html")
	//student
	render.AddFromFiles("addStudent.html", "templates/base/base.html", "templates/student/addStudent.html")

	return render

}
