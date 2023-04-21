package core

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	. "scm/db"
)

func GetLoginHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}
func Login(ctx *gin.Context) {
	//获取登录信息，账号、密码
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")
	//校验，获取数据库中的账号密码
	var userinfo UserInfo
	DB.Where("account = ? and password = ?", account, password).Take(&userinfo)
	if userinfo.ID == 0 {
		ctx.Redirect(http.StatusMovedPermanently, "/login")
	} else {
		/*//cookie验证
		//写cookie，登录成功设置cookie，键值对：isLogin：True
		ctx.SetCookie("isLogin", "true", 200, "/", "127.0.0.1", false, true)
		ctx.Redirect(http.StatusMovedPermanently, "/")*/

		//session验证
		session := sessions.Default(ctx)
		session.Set("isLogin", "true")
		session.Save()
		ctx.Redirect(http.StatusMovedPermanently, "/")
	}
}
