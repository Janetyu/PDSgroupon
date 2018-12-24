package orders

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
)

func ListForUser(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))

	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	userId, err := strconv.Atoi(c.DefaultQuery("uid", ""))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	infos, count, err := model.ListOrderModelByUser(offset, limit, userId)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "getlist from database has error: ")
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		OrdersList: infos,
	})
}

func ListForMerchant(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))

	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	merchantId, err := strconv.Atoi(c.DefaultQuery("mid", ""))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	infos, count, err := model.ListOrderModelByMerchant(offset, limit, merchantId)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "getlist from database has error: ")
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		OrdersList: infos,
	})
}

func ListForGoods(c *gin.Context) {
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))

	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	goodsId, err := strconv.Atoi(c.DefaultQuery("gid", ""))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	infos, count, err := model.ListOrderModelByGoods(offset, limit, goodsId)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "getlist from database has error: ")
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		OrdersList: infos,
	})
}
