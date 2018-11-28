package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"PDSgroupon/handler/sd"
	"PDSgroupon/handler/user"
	"PDSgroupon/router/middleware"
)

// 加载 中间件，路由器，处理器等
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
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

	u := g.Group("/v1/user")
	{
		u.POST("", user.Register)
		u.POST("/vcode", user.CreateVerifiCode)
		u.POST("/login", user.Login)
		u.POST("/loginbysms", user.LoginBySms)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
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
