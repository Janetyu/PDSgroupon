package permission

type CreateRequest struct {
	RoleName string `json:"role_name"`
}

type CreateResponse struct {
	Id uint64 `json:"id"`
	RoleName string `json:"role_name"`
}