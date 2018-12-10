package banner

import (
	"strconv"

	"github.com/gin-gonic/gin"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"os"
	"github.com/lexkong/log"
)

func Delete(c *gin.Context) {
	var filepath string

	bannerId, _ := strconv.Atoi(c.Param("id"))
	banner,err := model.GetBannerById(uint64(bannerId))
	if err != nil {
		SendResponse(c, errno.ErrBannerNotFount, nil)
		return
	}

	filepath = banner.Image

	if err := os.Remove(filepath); err != nil {
		SendResponse(c, errno.InternalServerError, nil)
		log.Errorf(err,"del file occured error is :")
		return
	}


	if err := model.DeleteBanner(uint64(bannerId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}