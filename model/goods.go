package model

import (
	"PDSgroupon/pkg/constvar"
	"gopkg.in/go-playground/validator.v9"
)

type GoodsModel struct {
	BaseModel
	GoodsName     string  `json:"goods_name" gorm:"column:goods_name"`         // 商品名称
	GoodsDesc     string  `json:"goods_desc" gorm:"column:goods_desc"`         // 商品描述
	GoodsPhoto    string  `json:"goods_photo" gorm:"column:goods_photo"`       // 商品图片
	GoodsCost     string `json:"goods_cost" gorm:"column:goods_cost"`         // 商品原价
	GoodsPrice    string `json:"goods_price" gorm:"column:goods_price"`       // 商品售价
	GoodsDiscount string `json:"goods_discount" gorm:"column:goods_discount"` // 拼团折扣价
	GoodsStock    int     `json:"goods_stock" gorm:"column:goods_stock"`       // 库存数
	StockWarn     int     `json:"stock_warn" gorm:"column:stock_warn"`         // 库存警告指标
	GoodsPeople   int     `json:"goods_people" gorm:"column:goods_people"`     // 拼团人数
	GroupAging    int     `json:"group_aging" gorm:"column:group_aging"`       // 拼团时长
	ShopId        uint64  `json:"shop_id" gorm:"column:shop_id"`               // 商铺id
	MainsortId    uint64  `json:"mainsort_id" gorm:"column:mainsort_id"`       // 主类别id
	SubsortId     uint64  `json:"subsort_id" gorm:"column:subsort_id"`         // 子类别id
	IsFare        bool    `json:"is_fare" gorm:"column:is_fare"`               //是否需要运费
	GoodsFare     string `json:"goods_fare" gorm:"column:goods_fare"`         // 商品运费
	GoodsSales	  int     `json:"goods_sales" gorm:"column:goods_sales"`	   // 商品销量
	IsShelf       bool    `json:"is_shelf" gorm:"column:is_shelf"`             // 是否上架
}

func (g *GoodsModel) TableName() string {
	return "goods"
}

func (g *GoodsModel) Create() error {
	return DB.Self.Create(&g).Error
}

func DeleteGoodsModel(id uint64) error {
	goods := GoodsModel{}
	goods.Id = id
	return DB.Self.Delete(&goods).Error
}

func (g *GoodsModel) Update() error {
	return DB.Self.Save(g).Error
}

func GetGoodsModel(name string) (*GoodsModel, error) {
	g := &GoodsModel{}
	d := DB.Self.Where("goods_name = ?", name).First(&g)
	return g, d.Error
}

func GetGoodsModelById(id uint64) (*GoodsModel, error) {
	g := &GoodsModel{}
	d := DB.Self.Where("id = ?", id).First(&g)
	return g, d.Error
}

func ListGoodsModel(offset, limit int) ([]*GoodsModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	goods := make([]*GoodsModel, 0)
	var count uint64

	if err := DB.Self.Model(&GoodsModel{}).Count(&count).Error; err != nil {
		return goods, count, err
	}

	if err := DB.Self.Offset(offset - 1).Limit(limit).Find(&goods).Error; err != nil {
		return goods, count, err
	}

	return goods, count, nil
}

func ListGoodsAll() ([]*GoodsModel, error) {

	goods := make([]*GoodsModel, 0)

	if err := DB.Self.Find(&goods).Error; err != nil {
		return goods, err
	}

	return goods, nil
}

func (g *GoodsModel) Validate() error {
	validate := validator.New()
	return validate.Struct(g)
}
