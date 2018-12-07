package admin

import (
	"github.com/gin-gonic/gin"

	. "PDSgroupon/handler"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/model"
	"PDSgroupon/pkg/token"
)

// 账号+密码登录
func AdminLogin(c *gin.Context) {
	var r LoginRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	admin, err := model.GetAdmin(r.Username)
	if err != nil {
		SendResponse(c, errno.ErrAdminNotFound, nil)
		return
	}

	err = admin.Compare(r.Password)
	if err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the json web token.
	t, err := token.Sign(c, token.Context{ID: admin.Id, Username: admin.Username,RoleId: admin.RoleId}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	rsp := LoginResponse{
		Id:        admin.Id,
		Username:  admin.Username,
		RoleId:    admin.RoleId,
		Token:     t,
	}

	SendResponse(c, nil, rsp)
}
