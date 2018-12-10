package category

import "PDSgroupon/model"

type CreateMainRequest struct {
	SortName string `json:"sort_name"`
}

type CreateSubRequest struct {
	Pid string `json:"pid"`
	SortName string `json:"sort_name"`
}

type CreateResponse struct {
	Id uint64 `json:"id"`
	Pid uint64 `json:"pid"`
	SortName string `json:"sort_name"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	CategoryList   []*model.CategoryModel `json:"categoryList"`
}