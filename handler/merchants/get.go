package merchants

import (
	"strconv"

	"github.com/gin-gonic/gin"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
)

func Get(c *gin.Context) {
	// merchantId
	id, _ := strconv.Atoi(c.Param("id"))
	merchant, err := model.GetMerchantById(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrMerchantNotFount, nil)
		return
	}

	SendResponse(c, nil, merchant)
}

// 检查用户是否已申请，如果有则直接返回数据，如果审核已通过，则前端引导重新登录
func MerchantStatus(c *gin.Context) {
	// userId
	uid, _ := strconv.Atoi(c.Param("uid"))

	mer, err := model.GetMerchantByOwnerId(uint64(uid))
	if err != nil && err.Error() != "record not found" {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, mer)
}
