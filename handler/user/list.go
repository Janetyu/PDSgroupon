package user

import (
	"github.com/gin-gonic/gin"

	. "PDSgroupon/handler"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/service"
	"strconv"
)

// List list the users in the database.
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

	infos, count, err := service.ListUser(offset, limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
