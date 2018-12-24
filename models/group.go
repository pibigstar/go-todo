package models

import (
	"time"

	"github.com/pibigstar/go-todo/models/db"
)

// MGroup 引用
var MGroup = &Group{}

// Group 组织表
type Group struct {
	ID            int
	GroupName     string    `gorm:"column:group_name"`
	GroupDescribe string    `gorm:"column:group_describe"`
	GroupMaster   string    `gorm:"column:group_master"`
	GroupCode     string    `gorm:"column:group_code"`
	CreateUser    int       `gorm:"column:create_user"`
	CreateTime    time.Time `gorm:"column:create_time"`
	UpdateTime    time.Time `gorm:"column:update_time"`
	IsDelete      bool      `gorm:"column:is_delete"`
}

// TableName 组织表
func (Group) TableName() string {
	return "groups"
}

// Create 创建
func (*Group) Create(group *Group) error {
	return db.Mysql.Create(&group)
}

// GetGroupByID 根据ID获取组织
func (group *Group) GetGroupByID(groupID int) (*Group, error) {
	err := db.Mysql.Where("id = ?", groupID).First(group).Error
	if err != nil {
		return nil, err
	}
	return group, nil
}

// GetGroupsByUserID 获取用户的所有组织
func (user *User) GetGroupsByUserID(userID int) ([]*Group, error) {
	var groups []*Group
	var groupsIds []int
	db.Mysql.Exec("select * from groups where id in(select group_id from group_user where user_id = ?)", userID)

	err := db.Mysql.Table("group_user").Where("user_id = ?", userID).Pluck("group_id", &groupsIds).Error
	if err != nil {
		return nil, err
	}

	err = db.Mysql.Table("groups").Where("id in (?)", groupsIds).Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return groups, nil
}
