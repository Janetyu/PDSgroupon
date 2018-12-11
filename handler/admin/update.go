package admin

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/util"
)

// 重置管理员密码
func AdminResetPwd(c *gin.Context) {
	log.Info("AdminResetPwd function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	// Get the user id from the url parameter.
	adminId, _ := strconv.Atoi(c.Param("id"))

	amodel, err := model.GetAdminById(uint64(adminId))
	if err != nil {
		SendResponse(c, errno.ErrAdminNotFound, nil)
		return
	}

	pwd := c.PostForm("password")
	newpwd := c.PostForm("newpassword")

	if err := checkPwdLen(pwd); err != nil {
		SendResponse(c, err, nil)
		return
	}

	if err := checkPwdLen(newpwd); err != nil {
		SendResponse(c, err, nil)
		return
	}

	if err := amodel.Compare(pwd); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	amodel.Password = newpwd

	if err := amodel.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err := amodel.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := GetOneResponse{
		Id:       amodel.Id,
		Username: amodel.Username,
		RoleId:   amodel.RoleId,
	}

	SendResponse(c, nil, rsp)
}
