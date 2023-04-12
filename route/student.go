package route

import (
	"github.com/gin-gonic/gin"
	. "scm/core"
)

func InitStuRoute(r *gin.RouterGroup) {
	r.GET("/", GetStudent)
	//get请求拿页面，POST请求拿数据
	r.GET("/add", GetAddStuHtml)
	r.POST("/add", AddStudent)
	r.GET("/delete/:delID", DeleteStudent)
	r.GET("/edit/:editID", GetEditStuHtml)
	r.POST("/edit/:editID", EditStudent)
}
