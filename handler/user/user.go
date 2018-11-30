package user

import (
	"PDSgroupon/model"
	"PDSgroupon/pkg/errno"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Vcode    string `json:"vcode"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginBySmsRequest struct {
	Username string `json:"username"`
	Vcode    string `json:"vcode"`
}

type LoginResponse struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	HeadImage string `json:"head_image"`
	RoleId    int64  `json:"role_id"`
	Token     string `json:"token"`
}

type GetOneResponse struct {
	Id        uint64  `json:"id"`
	Username  string  `json:"username"`
	NickName  string  `json:"nick_name"`
	Address   string  `json:"address"`
	Name      string  `json:"name"`
	HeadImage string  `json:"head_image"`
	Sex       string  `json:"sex"`
	Account   float64 `json:"account"`
	RoleId    int64   `json:"role_id"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}

func (r *CreateRequest) checkParam() error {
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
	}

	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}

	return nil
}

func checkPwdLen(pwd string) error {
	length := len(pwd)
	if length < 6 || length > 16 {
		return errno.New(errno.ErrValidation, nil).Add("password isn't enough len.")
	}
	return nil
}
