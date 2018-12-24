package models

import (
	"time"

	"github.com/pibigstar/go-todo/models/db"
)

// MUser 引用
var MUser = &User{}

// User 用户表
type User struct {
	ID         int       `gorm:"column:id"`
	OpenID     string    `gorm:"column:openId"`
	Phone      string    `gorm:"column:phone"`
	Gender     int       `gorm:"column:gender"`
	Password   string    `gorm:"column:password"`
	NickName   string    `gorm:"colunm:nick_name"`
	RealName   string    `gorm:"column:real_name"`
	AvatarURL  string    `gorm:"column:avatar_url"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

// TableName 用户表
func (User) TableName() string {
	return "users"
}

func (user *User) Create() error {
	return db.Mysql.Create(&user)
}

func (user *User) GetUserByOpenID(openID string) (*User, error) {
	err := db.Mysql.Where("opendId = ?", openID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
