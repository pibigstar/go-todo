package models

import (
	"time"

	"github.com/pibigstar/go-todo/models/db"
)

// MGroupUser 实体外部引用
var MGroupUser = &GroupUser{}

// GroupUser 组织用户关联实体
type GroupUser struct {
	ID         int       `gorm:"column:id"`
	GroupID    int       `gorm:"group_id"`
	UserID     int       `gorm:"user_id"`
	CreateTime time.Time `gorm:"column:create_time"`
	IsDelete   bool      `gorm:"column:is_delete"`
}

// Create 创建
func (*GroupUser) Create(groupUser *GroupUser) error {
	return db.Mysql.Create(groupUser)
}
