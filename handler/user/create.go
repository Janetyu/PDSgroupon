package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/util"
)

// Create creates a new user account.
func Register(c *gin.Context) {
	// X-Request-Id ,X-Correlation-Id 标识一个客户端和服务端的请求
	c.Set("X-Request-Id", util.UniqueId())
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	val, err := model.RC.GetKeyInRc(r.Vcode)
	if err != nil {
		SendResponse(c, errno.ErrVcodeNotFound, nil)
		return
	}

	log.Infof("the vcode get in redis is %s", val)

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
		RoleId:   1,
	}

	// 自定义验证函数
	//if err := r.checkParam(); err != nil {
	//	SendResponse(c, err, nil)
	//	return
	//}

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if user, err := model.GetUser(u.Username); err != nil || user.Username != "" {
		if user.Username != "" {
			SendResponse(c, errno.ErrUserHasRegist, nil)
			return
		} else if err != nil && err.Error() != "record not found" {
			SendResponse(c, errno.ErrDatabase, nil)
			log.Errorf(err, "the database error is:")
			return
		}
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	if err := model.RC.DelKeyInRc(r.Vcode); err != nil {
		log.Errorf(err, "The redis occurred error while del key: %s")
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}

// 创建并发送验证码，保存到redis数据库中
func CreateVerifiCode(c *gin.Context) {
	c.Set("X-Request-Id", util.UniqueId())
	log.Info("CreateVerifiCode function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	var p PhoneRequest
	if err := c.Bind(&p); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	pn := model.UserPhone{
		PhoneNum: p.PhoneNum,
	}

	if err := pn.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	vcode := util.GenerateVerificateCode()

	if err := RequestSms(pn.PhoneNum, vcode); err != nil {
		SendResponse(c, errno.ErrRequestSms, nil)
		return
	}

	if err := model.RC.SetKeyInRc(vcode, "120", vcode); err != nil {
		SendResponse(c, errno.ErrRedisConn, nil)
		log.Errorf(err, "Redis occurred error: %s")
		return
	}

	log.Infof("the vcode set in redis is %s", vcode)
	SendResponse(c, nil, nil)
}
