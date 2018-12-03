package model

import (
	"sync"
	"time"
)

type BaseModel struct {
	Id        uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"-"`
}

type UserInfo2 struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayHello"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
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

type UserList2 struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo2
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}
