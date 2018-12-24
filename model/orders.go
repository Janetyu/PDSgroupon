package model

import (
	"PDSgroupon/pkg/constvar"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type OrderModel struct {
	BaseModel
	OrderNum    string    `json:"order_num" gorm:"column:order_num"`
	MerchantId  uint64    `json:"merchant_id" gorm:"column:merchant_id"`
	GoodsId     uint64    `json:"goods_id" gorm:"column:goods_id"`
	ClientId    uint64    `json:"client_id" gorm:"column:client_id"`
	ClientNick  string    `json:"client_nick" gorm:"column:client_nick"`
	ClientPhone string    `json:"client_phone" gorm:"column:client_phone"`
	GroupId     uint64    `json:"group_id" gorm:"column:group_id"`
	OrderAddr   string    `json:"order_addr" gorm:"column:order_addr"`
	OrderPrice  string    `json:"order_price" gorm:"column:order_price"`
	OrderStatus string    `json:"order_status" gorm:"column:order_status"`
	OrderMark   string    `json:"order_mark" gorm:"column:order_mark"`
	PayMethod   string    `json:"pay_method" gorm:"cplumn:pay_method"`
	PayedAt     time.Time `gorm:"column:payedAt" json:"payedAt"`
	FinishedAt  time.Time `gorm:"column:finishedAt" json:"finishedAt"`
}

func (o *OrderModel) TableName() string {
	return "orders"
}

func (o *OrderModel) Create() error {
	// 事务控制
	tx := DB.Self.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&o).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func DeleteOrderModel(id uint64) error {
	orders := OrderModel{}
	orders.Id = id
	return DB.Self.Delete(&orders).Error
}

func (o *OrderModel) Update() error {
	return DB.Self.Save(o).Error
}

func GetOrderModel(orderNum string) (*OrderModel, error) {
	o := &OrderModel{}
	d := DB.Self.Where("order_num = ?", orderNum).First(&o)
	return o, d.Error
}

func GetOrderModelById(id uint64) (*OrderModel, error) {
	o := &OrderModel{}
	d := DB.Self.Where("id = ?", id).First(&o)
	return o, d.Error
}

func ListOrderModelByUser(offset, limit, clientId int) ([]*OrderModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	orders := make([]*OrderModel, 0)
	var count uint64

	if err := DB.Self.Model(&OrderModel{}).Where("client_id = ?", clientId).Count(&count).Error; err != nil {
		return orders, count, err
	}

	if err := DB.Self.Where("client_id = ?", clientId).Offset((offset - 1) * limit).Limit(limit).Find(&orders).Error; err != nil {
		return orders, count, err
	}

	return orders, count, nil
}

func ListOrderModelByMerchant(offset, limit, merchantId int) ([]*OrderModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	orders := make([]*OrderModel, 0)
	var count uint64

	if err := DB.Self.Model(&OrderModel{}).Where("merchant_id = ?", merchantId).Count(&count).Error; err != nil {
		return orders, count, err
	}

	if err := DB.Self.Where("merchant_id = ?", merchantId).Offset((offset - 1) * limit).Limit(limit).Find(&orders).Error; err != nil {
		return orders, count, err
	}

	return orders, count, nil
}

func ListOrderModelByGoods(offset, limit, goodsId int) ([]*OrderModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	orders := make([]*OrderModel, 0)
	var count uint64

	if err := DB.Self.Model(&OrderModel{}).Where("goods_id = ?", goodsId).Count(&count).Error; err != nil {
		return orders, count, err
	}

	if err := DB.Self.Where("goods_id = ?", goodsId).Offset((offset - 1) * limit).Limit(limit).Find(&orders).Error; err != nil {
		return orders, count, err
	}

	return orders, count, nil
}

func ListOrderAll() ([]*OrderModel, error) {

	orders := make([]*OrderModel, 0)

	if err := DB.Self.Find(&orders).Error; err != nil {
		return orders, err
	}

	return orders, nil
}

func (o *OrderModel) Validate() error {
	validate := validator.New()
	return validate.Struct(o)
}
