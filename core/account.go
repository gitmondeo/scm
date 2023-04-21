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
