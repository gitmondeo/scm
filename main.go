package main

import (
	"github.com/gin-gonic/gin"
	. "scm/db"     //导入数据库包
	. "scm/render" //导入第三方模板
	. "scm/route"  //导入路由
)

func main() {
	//获取gin引擎对象
	r := gin.Default()
	//加载静态文件
	r.Static("/static", "./static")
	//初始化数据库，连接数据库，迁移模型类生成SQL，生成表
	InitDB()
	//gin自带的加载模板
	//r.LoadHTMLFiles("templates/class.html")
	//使用第三方的模板加载文件
	r.HTMLRender = CreateMyRender()
	//路由
	InitRoute(r)

	r.Run()
}
