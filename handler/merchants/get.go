package merchants

import (
	"strconv"

	"github.com/gin-gonic/gin"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
)

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// Get the user by the `username` from the database.
	merchant, err := model.GetMerchantById(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrBannerNotFount, nil)
		return
	}

	rsp := merchant

	SendResponse(c, nil, rsp)
}