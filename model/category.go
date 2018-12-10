package model

import (
	"PDSgroupon/pkg/constvar"
	"gopkg.in/go-playground/validator.v9"
)

type CategoryModel struct {
	Id        uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id" `
	Pid       uint64    `json:"pid" gorm:"column:pid" `
	SortName string 	`json:"sort" gorm:"column:sort_name" `
}

func (c *CategoryModel)TableName() string {
	return "categorys"
}

func (c *CategoryModel)Create() error {
	return DB.Self.Create(&c).Error
}

func DeleteMainCategory(id uint64) error {
	var err error

	maincategory := CategoryModel{}
	maincategory.Id = id

	err = DB.Self.Delete(&maincategory).Error
	if err != nil {
		return err
	}

	err = DB.Self.Delete(CategoryModel{}, "pid = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteSubCategory(id uint64) error {
	category := CategoryModel{}
	category.Id = id
	return DB.Self.Delete(&category).Error
}

func (c *CategoryModel)Update() error {
	return DB.Self.Save(c).Error
}

func GetCategory(name string) (*CategoryModel, error) {
	c := &CategoryModel{}
	d := DB.Self.Where("sort_name = ?", name).First(&c)
	return c,d.Error
}

func GetCategoryById(id uint64) (*CategoryModel, error) {
	c := &CategoryModel{}
	d := DB.Self.Where("id = ?", id).First(&c)
	return c,d.Error
}

func ListMainCategory(offset, limit int) ([]*CategoryModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	categorys := make([]*CategoryModel,0)
	var count uint64

	if err := DB.Self.Model(&CategoryModel{}).Where("pid = ?",0).Count(&count).Error; err != nil {
		return categorys,count,err
	}

	if err := DB.Self.Where("pid = ?",0).Offset(offset - 1).Limit(limit).Order("id asc").Find(&categorys).Error;err != nil{
		return categorys,count,err
	}

	return categorys,count,nil
}

func ListSubCategory(offset, limit, pid int) ([]*CategoryModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	categorys := make([]*CategoryModel,0)
	var count uint64

	if err := DB.Self.Model(&CategoryModel{}).Where("pid = ?",pid).Count(&count).Error; err != nil {
		return categorys,count,err
	}

	if err := DB.Self.Where("pid = ?",pid).Offset(offset - 1).Limit(limit).Order("id asc").Find(&categorys).Error;err != nil{
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