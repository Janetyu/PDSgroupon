package goods

import (
	"os"
	"strconv"

	"github.com/lexkong/log"
	"github.com/gin-gonic/gin"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
)

func Delete(c *gin.Context) {
	var filepath string

	gId, _ := strconv.Atoi(c.Param("id"))
	goods, err := model.GetGoodsModelById(uint64(gId))
	if err != nil {
		SendResponse(c, errno.ErrGoodsNotFount, nil)
		return
	}

	filepath = goods.GoodsPhoto

	if err := os.Remove(filepath); err != nil {
		log.Errorf(err, "del shoplogo occured error is :")
	}

	if err := model.DeleteGoodsModel(uint64(gId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
