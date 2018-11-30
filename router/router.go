package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"PDSgroupon/handler/sd"
	"PDSgroupon/handler/upload"
	"PDSgroupon/handler/user"
	"PDSgroupon/router/middleware"
)

// 加载 中间件，路由器，处理器等
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 为 multipart 表单设置一个较低的内存限制（默认是 32 MiB）
	g.MaxMultipartMemory = 2 << 20 // 2 MB

	g.Static("/static", "./static")
	g.StaticFS("/more_static", http.Dir("static"))
	//g.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// 中间件函数
	g.Use(gin.Recovery())     // 用于panic时恢复API服务器
	g.Use(middleware.NoCache) // 强制浏览器不使用缓存
	g.Use(middleware.Options) // 浏览器跨域 OPTIONS 请求设置
	g.Use(middleware.Secure)  // 一些安全设置
	g.Use(mw...)

	// 404 的处理器函数
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	global := g.Group("/v1/global")
	{
		global.POST("/userlogin", user.UserLogin)
		global.POST("/userloginbysms", user.LoginBySms)
		global.POST("/vcode", user.CreateVerifiCode)
		global.POST("/register", user.Register)
	}

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.DELETE("/:id", user.Delete)
		u.PUT("/update/:id", user.Update)
		u.PUT("/resetpwd/:id", user.ResetPwd)
		u.PUT("/upload/:id", upload.SingleUpload)
		u.GET("/", user.List)
		u.GET("/:id", user.Get)
	}

	// 健康检查处理器的路由组
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
