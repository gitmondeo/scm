package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLoginHtml(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "login.html", nil)
}
