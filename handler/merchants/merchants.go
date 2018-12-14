package merchants

import (
	"time"
	
	"PDSgroupon/model"
)

type ReViewRequest struct {
	IsReview string `json:"is_review"`
	Mark string `json:"mark"` // 审核不通过备注
}

type GetOneResponse struct {
	Id uint64 `json:"id"`
	ShopName string `json:"shop_name"`
	ShopPhone string `json:"shop_phone"`
	ShopCert string `json:"shop_cert"` // 营业执照
	ShopQQ string `json:"shop_qq"`
	ShopLogo string `json:"shop_logo"`
	ShopIntro string `json:"shop_intro"`
	ShopAddr string `json:"shop_addr"`
	UserCert string `json:"owner_cert"` // 身份证
	UserId uint64 `json:"owner_id"`
	IsReview string `json:"is_review"`
	Mark string `json:"mark"`
	CreatedAt time.Time `json:"createdAt"`
}

type ListResponse struct {
	TotalCount uint64               `json:"totalCount"`
	MerchantList []*model.MerchantModel `json:"merchantList"`
}

