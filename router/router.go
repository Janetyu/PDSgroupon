package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"PDSgroupon/handler/admin"
	"PDSgroupon/handler/admin/banner"
	"PDSgroupon/handler/admin/category"
	"PDSgroupon/handler/admin/permission"
	"PDSgroupon/handler/goods"
	"PDSgroupon/handler/merchants"
	"PDSgroupon/handler/orders"
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
		global.POST("/register", user.Register)
		global.POST("/userlogin", user.UserLogin)
		global.POST("/userloginbysms", user.LoginBySms)
		global.POST("/vcode", user.CreateVerifiCode)

		global.POST("/adminlogin", admin.AdminLogin)

		global.GET("/mainsort/", category.MainList)
		global.GET("/subsort/", category.SubList)
		global.GET("/banner/", banner.List)

		global.GET("/goodsforhome/",goods.ListGoodsForHome)
		global.GET("/goodsforquery/",goods.ListGoodsForQuery)
	}

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.PUT("/update/:id", user.Update)
		u.PUT("/resetpwd/:id", user.ResetPwd)
		u.PUT("/upload/:id", upload.SingleUpload)
		u.GET("/detail/:id", user.Get)

		u.POST("/merchant/", merchants.Create)
		u.PUT("/merchant/:uid", merchants.UpdateForApply)
		u.GET("/merchant/:uid", merchants.MerchantStatus)

		u.POST("/orders/", orders.Create)
		u.GET("/orders/:id", orders.Get)
		u.GET("/orderlistbyuser/", orders.ListForUser)
	}

	m := g.Group("/v1/merchant")
	m.Use(middleware.AuthMiddleware())
	{
		m.GET("/detailbyuser/:uid", merchants.MerchantStatus)
		m.GET("/detailbyself/:id", merchants.Get)
		m.GET("/mainsort/subcount", category.MainListWithSubCount)
		m.GET("/mainsort/", category.MainList)
		m.GET("/subsort/", category.SubList)
		m.GET("/mainsortall/",category.MainListAll)
		m.GET("/subsortall/",category.SubListAllByMainId)
		m.GET("/suball/",category.SubListAll)
		m.PUT("/detail/:id", merchants.Update)

		m.POST("/goods/", goods.Create)
		m.PUT("/goods/:id", goods.Update)
		m.PUT("/goodsforshelf/:id", goods.IsShelf)
		m.GET("/goods/:id", goods.Get)
		m.GET("/goods/", goods.ListForMerchant)
		m.GET("/goodsbysub",goods.ListForMerchantAndSubSort)
		m.DELETE("/goods/:id", goods.Delete)
		m.DELETE("/goodsbymain/",goods.DeleteByMain)
		m.DELETE("/goodsbysub/",goods.DeleteBySub)

		m.GET("/orders/:id", orders.Get)
		m.GET("/orderlistbymerchant", orders.ListForMerchant)
		m.GET("/orderlistbygoods", orders.ListForGoods)
		m.PUT("/orders/:id", orders.UpdateStatus)
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

		a.POST("/merchant/", merchants.Create)
		a.PUT("/merchant/:id", merchants.Review)
		a.GET("/merchant/:id", merchants.Get)
		a.GET("/merchant/", merchants.List)
		a.DELETE("/merchant/:id", merchants.Delete)
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
