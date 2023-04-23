package route

import (
	"github.com/gin-gonic/gin"
	. "scm/core"
)

func InitAccountRoute(r *gin.RouterGroup) {
	r.GET("/", GetAccount)
	r.GET("/add", GetAddAccountHtml)
	r.POST("/add", AddAccount)
	r.GET("/edit/:editID", GetEditAccountHtml)
	r.POST("/edit/:editID", EditAccount)
	r.GET("/delete/:delID", DeleteAccount)
}
