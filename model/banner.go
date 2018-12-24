package model

import (
	"PDSgroupon/pkg/constvar"
	"gopkg.in/go-playground/validator.v9"
)

type BannerModel struct {
	BaseModel
	Title  string `json:"title" gorm:"column:title"`
	Url    string `json:"url" gorm:"column:url"`
	Image  string `json:"image" gorm:"column:image;not null" binding:"required" `
	Order  int    `json:"order" gorm:"column:order;unique"`
	CliNum int    `json:"cli_num" gorm:"column:cli_num"`
}

func (b *BannerModel) TableName() string {
	return "banners"
}

func (b *BannerModel) Create() error {
	return DB.Self.Create(&b).Error
}

func DeleteBanner(id uint64) error {
	banner := BannerModel{}
	banner.BaseModel.Id = id
	return DB.Self.Delete(&banner).Error
}

func (b *BannerModel) Update() error {
	return DB.Self.Save(b).Error
}

func GetBannerById(id uint64) (*BannerModel, error) {
	b := &BannerModel{}
	d := DB.Self.Where("id = ?", id).First(&b)
	return b, d.Error
}

func ListBanner(offset, limit int) ([]*BannerModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	banners := make([]*BannerModel, 0)
	var count uint64

	if err := DB.Self.Model(&BannerModel{}).Count(&count).Error; err != nil {
		return banners, count, err
	}

	if err := DB.Self.Offset((offset - 1) * limit).Limit(limit).Order("order").Find(&banners).Error; err != nil {
		return banners, count, err
	}

	return banners, count, nil
}

func ListBannerAll() ([]*BannerModel, error) {

	banners := make([]*BannerModel, 0)

	if err := DB.Self.Order("index asc").Find(&banners).Error; err != nil {
		return banners, err
	}

	return banners, nil
}

func (b *BannerModel) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}
