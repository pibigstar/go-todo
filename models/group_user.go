package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/pibigstar/go-todo/models/db"
)

// MGroupUser 实体外部引用
var MGroupUser = &GroupUser{}

// GroupUser 组织用户关联实体
type GroupUser struct {
	ID         int       `gorm:"column:id"`
	GroupID    int       `gorm:"group_id"`
	UserID     string    `gorm:"user_id"` //用户OpenID
	CreateTime time.Time `gorm:"column:create_time"`
	IsDelete   bool      `gorm:"column:is_delete"`
}

func (*GroupUser) Name() string {
	return "group_user"
}

// Create 创建
func (*GroupUser) Create(groupUser *GroupUser) error {
	return db.Mysql.Create(groupUser)
}

// GetGroupsByUserOpenID 获取用户加入的群
func (*GroupUser) GetUserJoinGroups(openID string) (*[]Group, error) {
	var groupIDs []int
	var groups []Group
	err := db.Mysql.Table("group_user").Where("user_id = ?", openID).Pluck("group_id", &groupIDs).Error
	if err == gorm.ErrRecordNotFound || len(groupIDs) == 0 {
		return nil, errors.New("此用户没有加入任何群")
	}
	err = db.Mysql.Find(&groups, "id in(?) and is_delete = ?", groupIDs, false).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("没有找到群")
	}
	return &groups, nil
}
