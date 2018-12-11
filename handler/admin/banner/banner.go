package banner

import (
	"PDSgroupon/model"
	"time"
)

type GetOneResponse struct {
	Id        uint64    `json:"id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	Order     int       `json:"order"`
	Image     string    `json:"image"`
	CliNum    int       `json:"cli_num"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ListResponse struct {
	TotalCount uint64               `json:"totalCount"`
	BannerList []*model.BannerModel `json:"bannerList"`
}
