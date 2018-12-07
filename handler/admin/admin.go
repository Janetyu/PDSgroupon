package admin

import (
	"PDSgroupon/pkg/errno"
	"PDSgroupon/model"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	RoleId    int64  `json:"role_id"`
	Token     string `json:"token"`
}

type GetOneResponse struct {
	Id        uint64  `json:"id"`
	Username  string  `json:"username"`
	RoleId    int64   `json:"role_id"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	AdminList   []*model.AdminInfo `json:"adminList"`
}

func checkPwdLen(pwd string) error {
	length := len(pwd)
	if length < 6 || length > 16 {
		return errno.New(errno.ErrValidation, nil).Add("password isn't enough len.")
	}
	return nil
}
