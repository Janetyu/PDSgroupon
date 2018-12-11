package model

import (
	"PDSgroupon/pkg/auth"
	"PDSgroupon/pkg/constvar"
	"gopkg.in/go-playground/validator.v9"
	"sync"
)

type AdminModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;unique;not null" binding:"required" validate:"min=5"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=6,max=128"`
	RoleId   int64  `json:"role_id" gorm:"column:role_id"`
}

type AdminInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	RoleId    int64  `json:"role_id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type AdminList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*AdminInfo
}

func (c *AdminModel) TableName() string {
	return "admins"
}

// Create creates a new user account.
func (a *AdminModel) Create() error {
	return DB.Self.Create(&a).Error
}

// DeleteUser deletes the user by the user identifier.
func DeleteAdmin(id uint64) error {
	admin := AdminModel{}
	admin.BaseModel.Id = id
	return DB.Self.Delete(&admin).Error
}

// Update updates an user account information.
func (a *AdminModel) Update() error {
	return DB.Self.Save(a).Error
}

// GetUser gets an user by the user identifier.
func GetAdmin(username string) (*AdminModel, error) {
	a := &AdminModel{}
	d := DB.Self.Where("username = ?", username).First(&a)
	return a, d.Error
}

func GetAdminById(id uint64) (*AdminModel, error) {
	a := &AdminModel{}
	d := DB.Self.Where("id = ?", id).First(&a)
	return a, d.Error
}

// ListUser List all users
func ListAdmin(offset, limit int) ([]*AdminModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	admins := make([]*AdminModel, 0)
	var count uint64

	if err := DB.Self.Model(&AdminModel{}).Count(&count).Error; err != nil {
		return admins, count, err
	}

	if err := DB.Self.Offset(offset - 1).Limit(limit).Order("id desc").Find(&admins).Error; err != nil {
		return admins, count, err
	}

	return admins, count, nil
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (a *AdminModel) Compare(pwd string) (err error) {
	err = auth.Compare(a.Password, pwd)
	return
}

// Encrypt the user password.
func (a *AdminModel) Encrypt() (err error) {
	a.Password, err = auth.Encrypt(a.Password)
	return
}

// Validate the fields.
func (a *AdminModel) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
