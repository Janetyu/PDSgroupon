package model

import (
	"PDSgroupon/pkg/constvar"
	"gopkg.in/go-playground/validator.v9"
)

type CategoryModel struct {
	BaseModel
	SortName string `json:"sort" gorm:"column:sort_name" `
}

func (c *CategoryModel)TableName() string {
	return "categorys"
}

func (c *CategoryModel)Create() error {
	return DB.Self.Create(&c).Error
}

func DeleteCategory(id uint64) error {
	category := CategoryModel{}
	category.BaseModel.Id = id
	return DB.Self.Delete(&category).Error
}

func (c *CategoryModel)Update() error {
	return DB.Self.Save(c).Error
}

func GetCategoryById(id uint64) (*CategoryModel, error) {
	c := &CategoryModel{}
	d := DB.Self.Where("id = ?", id).First(&c)
	return c,d.Error
}

func ListCategory(offset, limit int) ([]*CategoryModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	categorys := make([]*CategoryModel,0)
	var count uint64

	if err := DB.Self.Model(&CategoryModel{}).Count(&count).Error; err != nil {
		return categorys,count,err
	}

	if err := DB.Self.Offset(offset - 1).Limit(limit).Order("id asc").Find(&categorys).Error;err != nil{
		return categorys,count,err
	}

	return categorys,count,nil
}

func ListCategoryAll() ([]*CategoryModel, error) {

	categorys := make([]*CategoryModel,0)

	if err := DB.Self.Find(&categorys).Error;err != nil{
		return categorys,err
	}

	return categorys,nil
}


func (b *CategoryModel) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}