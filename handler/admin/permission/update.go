package permission

import (
	"time"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log/lager"
	"github.com/lexkong/log"

	. "PDSgroupon/handler"
	"PDSgroupon/util"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
)

func Update(c *gin.Context) {
	log.Info("Permission Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	// Get the user id from the url parameter.
	roleId, _ := strconv.Atoi(c.Param("id"))

	role, err := model.GetPermissionById(uint64(roleId))
	if err != nil {
		SendResponse(c, errno.ErrRoleNoFound, nil)
		return
	}

	roleName := c.DefaultPostForm("role_name", role.RoleName)

	pmodel := model.PermissionModel{
		Id: role.Id,
		RoleName:  roleName,
		CreatedAt: role.CreatedAt,
		UpdatedAt: time.Time{},
	}

	// Save changed fields.
	if err := pmodel.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Id: pmodel.Id,
		RoleName:  pmodel.RoleName,
		CreatedAt: pmodel.CreatedAt,
		UpdatedAt: pmodel.UpdatedAt,
	}

	SendResponse(c, nil, rsp)
}
