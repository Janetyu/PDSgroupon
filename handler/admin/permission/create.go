package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log/lager"
	"github.com/lexkong/log"

	. "PDSgroupon/handler"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/model"
	"PDSgroupon/util"
)

func Create(c *gin.Context)  {
	log.Info("Admin Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	p := model.PermissionModel{
		RoleName: r.RoleName,
	}

	if permission, err := model.GetPermission(p.RoleName); err != nil || permission.RoleName != "" {
		if permission.RoleName != "" {
			SendResponse(c, errno.ErrRoleHasCreate, nil)
			return
		} else if err != nil && err.Error() != "record not found" {
			SendResponse(c, errno.ErrDatabase, nil)
			log.Errorf(err, "the database error is:")
			return
		}
	}

	// Insert the user to the database.
	if err := p.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "the database error is:")
		return
	}

	newP, _ := model.GetPermission(p.RoleName)
	rsp := CreateResponse{
		Id: newP.Id,
		RoleName: newP.RoleName,
		CreatedAt: newP.CreatedAt,
		UpdatedAt: newP.UpdatedAt,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
