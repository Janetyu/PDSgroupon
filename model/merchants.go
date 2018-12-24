package model

import (
	"PDSgroupon/pkg/constvar"
	"gopkg.in/go-playground/validator.v9"
)

// 商铺Model
type MerchantModel struct {
	BaseModel
	ShopName  string `json:"shop_name" gorm:"column:shop_name"`
	ShopPhone string `json:"shop_phone" gorm:"column:shop_phone"`
	ShopCert  string `json:"shop_cert" gorm:"column:shop_cert"` // 营业执照
	ShopQQ    string `json:"shop_qq" gorm:"column:shop_qq"`
	ShopLogo  string `json:"shop_logo" gorm:"column:shop_logo"`
	ShopIntro string `json:"shop_intro" gorm:"column:shop_intro"`
	ShopAddr  string `json:"shop_addr" gorm:"column:shop_addr"`
	UserCert  string `json:"owner_cert" gorm:"column:owner_cert"` // 身份证
	UserId    uint64 `json:"owner_id" gorm:"column:owner_id"`
	IsReview  string `json:"is_review" gorm:"column:is_review"` // 是否通过审核
	Mark      string `json:"mark" gorm:"column:mark"`
}

func (m *MerchantModel) TableName() string {
	return "shops"
}

func (m *MerchantModel) Create() error {
	return DB.Self.Create(&m).Error
}

func DeleteMerchants(id uint64) error {
	m := MerchantModel{}
	m.BaseModel.Id = id
	return DB.Self.Delete(&m).Error
}

func (m *MerchantModel) Update() error {
	return DB.Self.Save(m).Error
}

func GetMerchantById(id uint64) (*MerchantModel, error) {
	m := &MerchantModel{}
	d := DB.Self.Where("id = ?", id).First(&m)
	return m, d.Error
}

func GetMerchantByOwnerId(id uint64) (*MerchantModel, error) {
	m := &MerchantModel{}
	d := DB.Self.Where("owner_id = ?", id).First(&m)
	return m, d.Error
}

func ListMerchant(offset, limit int) ([]*MerchantModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	merchants := make([]*MerchantModel, 0)
	var count uint64

	if err := DB.Self.Model(&MerchantModel{}).Count(&count).Error; err != nil {
		return merchants, count, err
	}

	if err := DB.Self.Offset((offset - 1) * limit).Limit(limit).Find(&merchants).Error; err != nil {
		return merchants, count, err
	}

	return merchants, count, nil
}

func ListMerchantAll() ([]*MerchantModel, error) {

	merchants := make([]*MerchantModel, 0)

	if err := DB.Self.Find(&merchants).Error; err != nil {
		return merchants, err
	}

	return merchants, nil
}

func (m *MerchantModel) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
