package goods

import (
	"strconv"

	"github.com/gin-gonic/gin"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
)

func Get(c *gin.Context) {
	// goodsId
	id, _ := strconv.Atoi(c.Param("id"))
	goods, err := model.GetGoodsModelById(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrGoodsNotFount, nil)
		return
	}

	SendResponse(c, nil, goods)
}
