package route

import (
	"github.com/gin-gonic/gin"
	. "scm/core"
)

func InitStuRoute(r *gin.RouterGroup) {
	r.GET("/", GetStudent)
	r.GET("/add", GetAddStuHtml)
	r.POST("/add", AddStudent)
}
