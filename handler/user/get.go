package user

import (
	"strconv"

	"github.com/gin-gonic/gin"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
)

// Get gets an user by the user id.
func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// Get the user by the `username` from the database.
	user, err := model.GetUserById(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	rsp := GetOneResponse{
		Id:        user.Id,
		Username:  user.Username,
		NickName:  user.NickName,
		Address:   user.Address,
		Name:      user.Name,
		HeadImage: user.HeadImage,
		Sex:       user.Sex,
		Account:   user.Account,
		RoleId:    user.RoleId,
	}

	SendResponse(c, nil, rsp)
}
