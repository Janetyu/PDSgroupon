package model

import (
	"fmt"

	"PDSgroupon/pkg/auth"
	"PDSgroupon/pkg/constvar"

	"gopkg.in/go-playground/validator.v9"
	"sync"
)

// User represents a registered user.
type UserModel struct {
	BaseModel
	Username  string  `json:"username" gorm:"column:phonenum;unique;not null" binding:"required" validate:"min=11,max=11"`
	Password  string  `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=6,max=128"`
	NickName  string  `json:"nick_name" gorm:"column:nick_name"`
	Address   string  `json:"address" gorm:"column:address"`
	Name      string  `json:"name" gorm:"column:name"`
	HeadImage string  `json:"head_image" gorm:"column:head_image"`
	Sex       string  `json:"sex" gorm:"column:sex;default:'男'"`
	Account   float64 `json:"account" gorm:"column:account;default:0"`
	RoleId    int64   `json:"role_id" gorm:"column:role_id"`
}

type UserPhone struct {
	PhoneNum string `json:"phone_num" binding:"required" validate:"min=11,max=11"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string  `json:"username"`
	NickName  string  `json:"nick_name"`
	Address   string  `json:"address"`
	Name      string  `json:"name"`
	HeadImage string  `json:"head_image"`
	Sex       string  `json:"sex"`
	Account   float64 `json:"account"`
	RoleId    int64   `json:"role_id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

type UserInfo2 struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayHello"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserList2 struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo2
}

func (c *UserModel) TableName() string {
	return "users"
}

// Create creates a new user account.
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

// DeleteUser deletes the user by the user identifier.
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

// Update updates an user account information.
func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

// GetUser gets an user by the user identifier.
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("phonenum = ?", username).First(&u)
	return u, d.Error
}

func GetUserById(id uint64) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("id = ?", id).First(&u)
	return u, d.Error
}

// ListUser List all users
func ListUser(offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	if err := DB.Self.Model(&UserModel{}).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Offset(offset - 1).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// ListUser List all users
func ListUser2(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// Encrypt the user password.
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Validate the fields.
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (p *UserPhone) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
