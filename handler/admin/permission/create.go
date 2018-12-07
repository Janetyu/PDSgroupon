package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log/lager"
)

func Create(c *gin.Context)  {
	log.Info("Admin Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
}
