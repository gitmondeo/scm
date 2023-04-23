package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "scm/db"
)

func GetAccount(ctx *gin.Context) {
	var userInfos []UserInfo
	//获取模糊搜索参数
	searchParams := ctx.Query("searchParams")

	if searchParams == "" {
		DB.Find(&userInfos)
	} else {
		DB.Where("account like ?", "%"+searchParams+"%").Find(&userInfos)
	}

	ctx.HTML(http.StatusOK, "account.html", gin.H{
		"userInfos": userInfos,
	})
}
func GetAddAccountHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "addAccount.html", nil)
}
func AddAccount(ctx *gin.Context) {
	//获取前端页面数据,赋值到userinfo对象
	account := ctx.PostForm("account")
	pwd := ctx.PostForm("pwd")
	userinfo := UserInfo{Account: account, Password: pwd}
	//存到数据库
	DB.Create(&userinfo)
	ctx.Redirect(http.StatusMovedPermanently, "/account")
}

func GetEditAccountHtml(ctx *gin.Context) {
	//获取要编辑的账号的ID
	editID := ctx.Param("editID")
	//根据ID找到对应的账号信息
	var userinfo UserInfo
	DB.Where("id = ?", editID).Find(&userinfo)
	ctx.HTML(http.StatusOK, "editAccount.html", gin.H{
		"userinfo": userinfo,
	})
}
func EditAccount(ctx *gin.Context) {
	//获取前端页面数据,赋值到userinfo对象
	id := ctx.PostForm("id")
	account := ctx.PostForm("account")
	pwd := ctx.PostForm("pwd")
	DB.Where("id = ?", id).Updates(&UserInfo{Account: account, Password: pwd})
	ctx.Redirect(http.StatusMovedPermanently, "/account")
}
func DeleteAccount(ctx *gin.Context) {
	delID := ctx.Param("delID")
	DB.Where("id = ?", delID).Delete(&UserInfo{})
	ctx.Redirect(http.StatusMovedPermanently, "/account")
}
