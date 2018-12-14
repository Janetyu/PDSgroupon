package merchants

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
)

func Delete(c *gin.Context) {
	var filepath string

	mId, _ := strconv.Atoi(c.Param("id"))
	merchant, err := model.GetMerchantById(uint64(mId))
	if err != nil {
		SendResponse(c, errno.ErrMerchantNotFount, nil)
		return
	}

	uid := merchant.UserId
	user,err := model.GetUserById(uid)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	user.RoleId = 1
	if err := user.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	filepath = merchant.ShopLogo

	if err := os.Remove(filepath); err != nil {
		log.Errorf(err, "del shoplogo occured error is :")
	}

	if err := model.DeleteMerchants(uint64(mId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
