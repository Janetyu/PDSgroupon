package goods

import "PDSgroupon/model"

type ListResponse struct {
	TotalCount   uint64             `json:"totalCount"`
	GoodsList []*model.GoodsModel 	`json:"goodsList"`
}

