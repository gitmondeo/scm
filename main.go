package main

import (
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	. "scm/db"     //导入数据库包
	. "scm/render" //导入第三方模板
	. "scm/route"  //导入路由
	// 导入session包
	"github.com/gin-contrib/sessions"
)

func main() {
	//获取gin引擎对象
	r := gin.Default()

	//配置session,初始化session
	// 创建基于cookie的存储引擎，yuan 参数是用于加密的密钥，可以随便填写
	store := cookie.NewStore([]byte("yuan"))
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎
	r.Use(sessions.Sessions("mysession", store))

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

	r.Run("10.172.12.24:80")
}
