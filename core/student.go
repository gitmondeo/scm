package core

import "github.com/gin-gonic/gin"

func GetStudent(ctx *gin.Context) {
	ctx.HTML(200, "student.html", nil)
}
func GetAddStuHtml(ctx *gin.Context) {

	ctx.HTML(200, "addStudent.html", nil)
}
func AddStudent(ctx *gin.Context) {
	//获取前端请求数据

	//数据库存储
}
