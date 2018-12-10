package banner

import (
	"strconv"

	"github.com/gin-gonic/gin"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"github.com/lexkong/log"
)

func List(c *gin.Context) {
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

	infos, count, err := model.ListBanner(offset, limit)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err,"getlist from database has error: ")
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		BannerList:   infos,
	})
}
