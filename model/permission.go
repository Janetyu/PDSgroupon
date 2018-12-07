package model

import (
	"PDSgroupon/pkg/constvar"
	"gopkg.in/go-playground/validator.v9"
)

type PermissionModel struct {
	BaseModel
	RoleName string `json:"role" gorm:"column:"`
}

func (p *PermissionModel)TableName() string {
	return "roles"
}

func (p *PermissionModel)Create() error {
	return DB.Self.Create(&p).Error
}

func DeletePermission(id uint64) error {
	permission := PermissionModel{}
	permission.BaseModel.Id = id
	return DB.Self.Delete(&permission).Error
}

func (p *PermissionModel)Update() error {
	return DB.Self.Save(p).Error
}

func GetPermissionById(id uint64) (*PermissionModel, error) {
	p := &PermissionModel{}
	d := DB.Self.Where("id = ?", id).First(&p)
	return p,d.Error
}

func ListPermission(offset, limit int) ([]*PermissionModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	permissions := make([]*PermissionModel,0)
	var count uint64

	if err := DB.Self.Model(&PermissionModel{}).Count(&count).Error; err != nil {
		return permissions,count,err
	}

	if err := DB.Self.Offset(offset - 1).Limit(limit).Order("id asc").Find(&permissions).Error;err != nil{
		return permissions,count,err
	}

	return permissions,count,nil
}

func ListPermissionAll() ([]*PermissionModel, error) {

	permissions := make([]*PermissionModel,0)

	if err := DB.Self.Find(&permissions).Error;err != nil{
		return permissions,err
	}

	return permissions,nil
}


func (p *PermissionModel) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}