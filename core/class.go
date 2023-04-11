package core

import (
	"github.com/gin-gonic/gin"
	. "scm/db"
	"strconv"
	//. "scm/db"
)

func GetClass(ctx *gin.Context) {
	ctx.HTML(200, "class.html", nil)
}
func GetClassHtml(ctx *gin.Context) {
	ctx.HTML(200, "addClass.html", nil)

}

func AddClass(ctx *gin.Context) {
	name := ctx.PostForm("name")
	num, _ := strconv.Atoi(ctx.PostForm("num"))
	tutor, _ := strconv.Atoi(ctx.PostForm("tutor"))
	addClass := Class{Base: Base{Name: name}, Num: num, TutorID: tutor}
	DB.Create(&addClass)

	var classes []Class
	DB.Find(&classes)
	ctx.JSON(200, gin.H{
		"classes": classes,
	})

}
