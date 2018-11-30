package middleware

import (
	"github.com/gin-gonic/gin"

	"PDSgroupon/handler"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			// 假设您有一个授权中间件验证当前请求是否已授权。
			// 如果授权失败（例如：密码不匹配），请调用Abort以确保剩余的处理程序
			c.Abort()
			return
		}

		c.Next()
	}
}
