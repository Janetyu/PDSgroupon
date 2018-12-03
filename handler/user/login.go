package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/pkg/token"
)

// 手机号+密码登录
func UserLogin(c *gin.Context) {
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

	// Sign the json web token.
	t, err := token.Sign(c, token.Context{ID: user.Id, Username: user.Username,RoleId: user.RoleId}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	rsp := LoginResponse{
		Id:        user.Id,
		Username:  user.Username,
		HeadImage: user.HeadImage,
		RoleId:    user.RoleId,
		Token:     t,
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

	// Sign the json web token.
	t, err := token.Sign(c, token.Context{ID: user.Id, Username: user.Username,RoleId: user.RoleId}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	rsp := LoginResponse{
		Id:        user.Id,
		Username:  user.Username,
		HeadImage: user.HeadImage,
		RoleId:    user.RoleId,
		Token:     t,
	}

	SendResponse(c, nil, rsp)
}
