package goods

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/service"
)

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

	mid, err := strconv.Atoi(c.DefaultQuery("mid", "0"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	infos, count, err := model.ListGoodsModelForMerchant(offset, limit,mid)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "getlist from database has error: ")
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		GoodsList:  infos,
	})
}

func ListForMerchantAndSubSort(c *gin.Context) {
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

	mid, err := strconv.Atoi(c.DefaultQuery("mid", "0"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	sid, err := strconv.Atoi(c.DefaultQuery("sid", "0"))
	if err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	infos, count, err := model.ListGoodsModelForMerchantAndSub(offset, limit, mid, sid)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "getlist from database has error: ")
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		GoodsList:  infos,
	})
}

func ListGoodsForHome(c *gin.Context) {

	infos, err := service.ListGoodsByAllMainSort()
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "getlist from database has error: ")
		return
	}

	SendResponse(c, nil, GoodsListWithMainsort{
		GoodsList:  infos,
	})
}

func ListGoodsForQuery(c *gin.Context)  {

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

	q := c.DefaultQuery("query", "")

	infos,count, err := model.ListGoodsBySearch(offset,limit,q)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "getlist from database has error: ")
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		GoodsList:  infos,
	})
}

