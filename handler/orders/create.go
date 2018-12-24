package orders

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	. "PDSgroupon/handler"
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
	"PDSgroupon/util"
	"time"
)

func Create(c *gin.Context) {
	log.Info("Order Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateOrder
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	//tm := time.Date(0, 0, 0, 0, 0, 0, 0, time.Now().Location())

	o := model.OrderModel{
		OrderNum:    util.UniqueId(),
		MerchantId:  r.MerchantId,
		GoodsId:     r.GoodsId,
		ClientId:    r.ClientId,
		ClientNick:  r.ClientNick,
		ClientPhone: r.ClientPhone,
		GroupId:     r.GroupId,
		OrderAddr:   r.OrderAddr,
		OrderPrice:  r.OrderPrice,
		OrderMark:   r.OrderMark,
		OrderStatus: "待发货",
		PayMethod:   r.PayMethod,
	}

	// Insert the user to the database.
	if err := o.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "the database error is:")
		return
	}

	mer, err := model.GetMerchantById(o.MerchantId)
	if err != nil {
		SendResponse(c, errno.ErrMerchantNotFount, nil)
		return
	}

	goods, err := model.GetGoodsModelById(o.GoodsId)
	if err != nil {
		SendResponse(c, errno.ErrGoodsNotFount, nil)
		return
	}

	rsp := CreateResponse{
		OrderNum:     o.OrderNum,
		MerchantName: mer.ShopName,
		Goods:        goods,
		ClientId:     o.ClientId,
		ClientNick:   o.ClientNick,
		ClientPhone:  o.ClientPhone,
		OrderAddr:    o.OrderAddr,
		OrderPrice:   o.OrderPrice,
		OrderMark:    o.OrderMark,
		PayMethod:    o.PayMethod,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		PayedAt:      o.PayedAt,
		FinishedAt:   o.FinishedAt,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
