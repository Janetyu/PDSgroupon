package user

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/util"
)

// Update update a exist user account info.
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	// Get the user id from the url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := model.GetUserById(uint64(userId))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	nickName := c.DefaultPostForm("nick_name", user.NickName)
	address := c.DefaultPostForm("address", user.Address)
	name := c.DefaultPostForm("name", user.Name)
	sex := c.DefaultPostForm("sex", user.Sex)

	umodel := model.UserModel{
		BaseModel: model.BaseModel{Id: user.Id, CreatedAt: user.CreatedAt, UpdatedAt: time.Time{}},
		Username:  user.Username,
		Password:  user.Password,
		NickName:  nickName,
		Address:   address,
		Name:      name,
		HeadImage: user.HeadImage,
		Sex:       sex,
		Account:   user.Account,
		RoleId:    user.RoleId,
	}

	// Validate the data.
	if err := umodel.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Save changed fields.
	if err := umodel.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := GetOneResponse{
		Id:        umodel.Id,
		Username:  umodel.Username,
		NickName:  umodel.NickName,
		Address:   umodel.Address,
		Name:      umodel.Name,
		HeadImage: umodel.HeadImage,
		Sex:       umodel.Sex,
		Account:   umodel.Account,
		RoleId:    umodel.RoleId,
	}

	SendResponse(c, nil, rsp)
}

// 重置密码
func ResetPwd(c *gin.Context) {
	log.Info("ResetPwd function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	// Get the user id from the url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	umodel, err := model.GetUserById(uint64(userId))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
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

	if err := umodel.Compare(pwd); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	umodel.Password = newpwd

	if err := umodel.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err := umodel.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := GetOneResponse{
		Id:        umodel.Id,
		Username:  umodel.Username,
		NickName:  umodel.NickName,
		Address:   umodel.Address,
		Name:      umodel.Name,
		HeadImage: umodel.HeadImage,
		Sex:       umodel.Sex,
		Account:   umodel.Account,
		RoleId:    umodel.RoleId,
	}

	SendResponse(c, nil, rsp)
}
