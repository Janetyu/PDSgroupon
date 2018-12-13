package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"PDSgroupon/handler/admin"
	"PDSgroupon/handler/admin/banner"
	"PDSgroupon/handler/admin/category"
	"PDSgroupon/handler/admin/permission"
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
		global.POST("/adminlogin", admin.AdminLogin)
		global.POST("/vcode", user.CreateVerifiCode)
		global.POST("/register", user.Register)
		global.GET("/mainsort/", category.MainList)
		global.GET("/subsort/", category.SubList)
		global.GET("/banner/", banner.List)
	}

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.PUT("/update/:id", user.Update)
		u.PUT("/resetpwd/:id", user.ResetPwd)
		u.PUT("/upload/:id", upload.SingleUpload)
		u.GET("/:id", user.Get)
	}

	a := g.Group("/v1/admin")
	a.Use(middleware.AuthMiddleware())
	{
		a.POST("/register", admin.Register)
		a.PUT("/detail/:id", admin.AdminResetPwd)
		a.GET("/detail/:id", admin.Get)
		a.GET("/", admin.List)

		a.GET("/userlist/", user.List)
		a.DELETE("/userdel/:id", user.Delete)
		a.PUT("/userupd/:id", user.Update)

		a.POST("/roleadd", permission.Create)
		a.GET("/rolelist", permission.List)
		a.PUT("/roleupd/:id", permission.Update)

		a.POST("/banner/", banner.Create)
		a.GET("/banner/:id", banner.Get)
		a.GET("/banner/", banner.List)
		a.DELETE("/banner/:id", banner.Delete)
		a.PUT("/banner/:id", banner.Update)
		a.PUT("/bannerupload/:id", upload.SingleUpload)

		a.POST("/mainsortadd/", category.CreateMain)
		a.POST("/subsortadd/", category.CreateSub)
		a.GET("/mainsort/", category.MainList)
		a.GET("/subsort/", category.SubList)
		a.PUT("/mainsort/:id", category.UpdateMain)
		a.PUT("/subsort/:id", category.UpdateSub)
		a.DELETE("/mainsort/:id", category.DeleteMain)
		a.DELETE("/subsort/:id", category.DeleteSub)
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
