package permission

import (
	"time"
	"PDSgroupon/model"
)

type CreateRequest struct {
	RoleName string `json:"role_name"`
}

type CreateResponse struct {
	Id uint64 `json:"id"`
	RoleName string `json:"role_name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	PermissionList   []*model.PermissionModel `json:"roleList"`
}