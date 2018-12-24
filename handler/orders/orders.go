package orders

import (
	"PDSgroupon/model"
	"time"
)

type CreateOrder struct {
	MerchantId  uint64 `json:"merchant_id"`
	GoodsId     uint64 `json:"goods_id"`
	ClientId    uint64 `json:"client_id"`
	ClientNick  string `json:"client_nick"`
	ClientPhone string `json:"client_phone"`
	GroupId     uint64 `json:"group_id"`
	OrderAddr   string `json:"order_addr"`
	OrderPrice  string `json:"order_price"`
	OrderMark   string `json:"order_mark"`
	PayMethod   string `json:"pay_method"`
}

type CreateResponse struct {
	OrderNum     string            `json:"order_num"`
	MerchantName string            `json:"merchant_name"`
	Goods        *model.GoodsModel `json:"goods"`
	ClientId     uint64            `json:"client_id"`
	ClientNick   string            `json:"client_nick"`
	ClientPhone  string            `json:"client_phone"`
	OrderAddr    string            `json:"order_addr"`
	OrderPrice   string            `json:"order_price"`
	OrderMark    string            `json:"order_mark"`
	PayMethod    string            `json:"pay_method"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
	PayedAt      time.Time         `json:"payedAt"`
	FinishedAt   time.Time         `json:"finishedAt"`
}

type UpdateRequest struct {
	OrderStatus string `json:"order_status"`
}

type ListResponse struct {
	TotalCount uint64              `json:"totalCount"`
	OrdersList []*model.OrderModel `json:"ordersList"`
}
