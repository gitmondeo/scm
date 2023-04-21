package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndex(ctx *gin.Context) {
	//cookie检验是否登录成功
	isLogin, _ := ctx.Cookie("isLogin")
	if isLogin == "true" {
		ctx.HTML(200, "base.html", nil)
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "/login")
	}

	//session判断是否登录成功
	/*session := sessions.Default(ctx)
	isLogin := session.Get("isLogin")
	if isLogin == "true" {
		ctx.HTML(200, "base.html", nil)
	} else {
		ctx.Redirect(301, "/login")
	}*/
}
