package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	. "scm/db"
	"strconv"
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
		//cookie验证
		//写cookie，登录成功设置cookie，键值对：isLogin：True
		ctx.SetCookie("isLogin", "true", 2000, "/", "10.172.12.24", false, true)
		fmt.Println("11111")
		fmt.Println(ctx.Cookie("isLogin"))
		//session验证
		/*session := sessions.Default(ctx)
		session.Set("isLogin", "true")
		session.Save()*/

		if userinfo.ID == 1 {
			//root账号
			//ctx.Redirect(http.StatusMovedPermanently, "/")
			ctx.HTML(http.StatusOK, "base.html", nil)

		} else {
			//查询学生信息，找出学生学号，学生登录跳转到个人界面
			var stu Student
			DB.Where("name = ?", userinfo.Account).Take(&stu)
			sno := strconv.Itoa(stu.Sno)
			ctx.Redirect(http.StatusMovedPermanently, "/student/"+sno)
		}

	}
}
