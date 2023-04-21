package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetIndex(ctx *gin.Context) {
	//cookie检验是否登录成功
	isLogin, _ := ctx.Cookie("isLogin")
	fmt.Println("isLogin111::::", isLogin)
	/*if isLogin == "true" {
		ctx.HTML(200, "base.html", nil)
	} else {
		fmt.Println("isLogin222::::", isLogin)
		ctx.Redirect(http.StatusMovedPermanently, "/login")
	}*/
	if isLogin != "nil" {
		ctx.HTML(200, "base.html", nil)
	} else {
		fmt.Println("hhhhhh")
		//ctx.Redirect(http.StatusMovedPermanently, "/login")
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
