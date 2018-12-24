package orders

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/util"
	"time"
)

func UpdateStatus(c *gin.Context) {
	log.Info("UpdateOrderStatus function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	var r UpdateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	o, err := model.GetOrderModelById(uint64(id))
	if err != nil {
		SendResponse(c, errno.ErrOrderNotFount, nil)
		return
	}

	if r.OrderStatus == "已支付" {
		o.PayedAt = time.Time{}
		o.OrderStatus = r.OrderStatus
	}

	if r.OrderStatus == "已发货" {
		o.UpdatedAt = time.Time{}
		o.OrderStatus = r.OrderStatus
	}

	if r.OrderStatus == "已完成" {
		o.FinishedAt = time.Time{}
		o.OrderStatus = r.OrderStatus
	}

	if err := o.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "the database error is:")
		return
	}

	SendResponse(c, nil, o)
}
