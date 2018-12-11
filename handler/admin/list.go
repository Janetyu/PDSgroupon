package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"

	. "PDSgroupon/handler"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/service"
)

func List(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))

	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	infos, count, err := service.ListAdmin(offset, limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		AdminList:  infos,
	})
}
