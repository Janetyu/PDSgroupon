package model

import (
	"PDSgroupon/pkg/constvar"
	"time"
)

type PermissionModel struct {
	Id        uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	RoleName  string    `json:"role_name" gorm:"column:role_name"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func (p *PermissionModel) TableName() string {
	return "roles"
}

func (p *PermissionModel) Create() error {
	return DB.Self.Create(&p).Error
}

func DeletePermission(id uint64) error {
	permission := PermissionModel{}
	permission.Id = id
	return DB.Self.Delete(&permission).Error
}

func (p *PermissionModel) Update() error {
	return DB.Self.Save(p).Error
}

func GetPermission(roleName string) (*PermissionModel, error) {
	p := &PermissionModel{}
	d := DB.Self.Where("role_name = ?", roleName).First(&p)
	return p, d.Error
}

func GetPermissionById(id uint64) (*PermissionModel, error) {
	p := &PermissionModel{}
	d := DB.Self.Where("id = ?", id).First(&p)
	return p, d.Error
}

func ListPermission(offset, limit int) ([]*PermissionModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	permissions := make([]*PermissionModel, 0)
	var count uint64

	if err := DB.Self.Model(&PermissionModel{}).Count(&count).Error; err != nil {
		return permissions, count, err
	}

	if err := DB.Self.Offset(offset - 1).Limit(limit).Order("id asc").Find(&permissions).Error; err != nil {
		return permissions, count, err
	}

	return permissions, count, nil
}

func ListPermissionAll() ([]*PermissionModel, error) {

	permissions := make([]*PermissionModel, 0)

	if err := DB.Self.Find(&permissions).Error; err != nil {
		return permissions, err
	}

	return permissions, nil
}
