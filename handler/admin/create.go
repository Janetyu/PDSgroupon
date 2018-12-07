package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log/lager"
	"github.com/lexkong/log"

	. "PDSgroupon/handler"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/util"
	"PDSgroupon/model"
)

func Register(c *gin.Context) {
	// X-Request-Id ,X-Correlation-Id 标识一个客户端和服务端的请求
	//c.Set("X-Request-Id", util.UniqueId())
	log.Info("Admin Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	a := model.AdminModel{
		Username: r.Username,
		Password: r.Password,
		RoleId:   3,
	}

	// Validate the data.
	if err := a.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if admin, err := model.GetAdmin(a.Username); err != nil || admin.Username != "" {
		if admin.Username != "" {
			SendResponse(c, errno.ErrAdminHasRegist, nil)
			return
		} else if err != nil && err.Error() != "record not found" {
			SendResponse(c, errno.ErrDatabase, nil)
			log.Errorf(err, "the database error is:")
			return
		}
	}

	// Encrypt the user password.
	if err := a.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := a.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "the database error is:")
		return
	}

	rsp := CreateResponse{
		Username: a.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
