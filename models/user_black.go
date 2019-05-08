package models

import (
	"github.com/pibigstar/go-todo/models/db"
	"time"
)

// MBlackUser 引用
var MBlackUser = &BlackUser{}

// User 用户表
type BlackUser struct {
	ID         int       `gorm:"column:id"`
	UserName   string    `gorm:"column:username"`
	Reason     string    `gorm:"column:reason"`
	OpenId     string    `gorm:"column:open_id"`
	IsDelete   bool      `gorm:"column:is_delete"`
	CreateTime time.Time `gorm:"column:create_time"`
}

// TableName 用户表
func (*BlackUser) TableName() string {
	return "user_black"
}

func (t *BlackUser) ListBlack() ([]BlackUser, error) {
	var admins []BlackUser
	err := db.Mysql.Table(t.TableName()).Where("is_delete = ?", false).Find(&admins).Error
	if err != nil {
		return nil, err
	}
	return admins, nil
}
