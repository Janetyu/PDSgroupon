package middleware

import (
	"github.com/gin-gonic/gin"

	"PDSgroupon/handler"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/pkg/token"
	"regexp"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx , err := token.ParseRequest(c)
		if  err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			// 假设您有一个授权中间件验证当前请求是否已授权。
			// 如果授权失败（例如：密码不匹配），请调用Abort以确保剩余的处理程序
			c.Abort()
			return
		}

		// 全局接口拦截器，实现api鉴权，原理是路径匹配

		roleid := ctx.RoleId
		path := c.Request.URL.Path
		user_reg := regexp.MustCompile("/v1/user")
		merchant_reg := regexp.MustCompile("/v1/merchant")
		admin_reg := regexp.MustCompile("/v1/admin")

		if roleid < 1 && user_reg.MatchString(path){
			handler.SendResponse(c, errno.ErrPermission, nil)
			c.Abort()
			return
		}
		if roleid != 2 && merchant_reg.MatchString(path){
			handler.SendResponse(c, errno.ErrPermission, nil)
			c.Abort()
			return
		}
		if roleid != 3 && admin_reg.MatchString(path){
			handler.SendResponse(c, errno.ErrPermission, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}

