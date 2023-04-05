package core

import "github.com/gin-gonic/gin"

func GetIndex(ctx *gin.Context) {
	ctx.HTML(200, "base.html", nil)
}
