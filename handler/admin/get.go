package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
)

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// Get the user by the `username` from the database.
	admin, err := model.GetAdminById(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrAdminNotFound, nil)
		return
	}

	rsp := GetOneResponse{
		Id:       admin.Id,
		Username: admin.Username,
		RoleId:   admin.RoleId,
	}

	SendResponse(c, nil, rsp)
}
