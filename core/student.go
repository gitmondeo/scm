package core

import "github.com/gin-gonic/gin"

func GetStudent(ctx *gin.Context) {
	ctx.HTML(200, "student.html", nil)
}
