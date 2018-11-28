package user

import (
	"github.com/gin-gonic/gin"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"github.com/lexkong/log"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}

// 手机号+密码登录
func Login(c *gin.Context) {
	var r LoginRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	// Get the user by the `username` from the database.
	user, err := model.GetUser(r.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	err = user.Compare(r.Password)
	if err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	rsp := CreateResponse{
		Username: user.Username,
	}

	SendResponse(c, nil, rsp)
}

// 手机+短信验证码
func LoginBySms(c *gin.Context) {
	var r LoginBySmsRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	_, err := model.RC.GetKeyInRc(r.Vcode)
	if err != nil {
		SendResponse(c, errno.ErrVcodeNotFound, nil)
		return
	}

	// Get the user by the `username` from the database.
	user, err := model.GetUser(r.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	if err := model.RC.DelKeyInRc(r.Vcode); err != nil {
		log.Errorf(err, "The redis occurred error while del key: %s")
	}

	rsp := CreateResponse{
		Username: user.Username,
	}

	SendResponse(c, nil, rsp)
}
