package orders

import (
	"strconv"

	"github.com/gin-gonic/gin"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
)

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	order, err := model.GetOrderModelById(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrOrderNotFount, nil)
		return
	}

	SendResponse(c, nil, order)
}
