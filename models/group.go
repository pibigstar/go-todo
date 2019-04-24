package models

import (
	"github.com/gogf/gf/g/util/gconv"
	"github.com/jinzhu/gorm"
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
	CreateUser    string    `gorm:"column:create_user"`
	CreateTime    time.Time `gorm:"column:create_time"`
	UpdateTime    time.Time `gorm:"column:update_time"`
	IsDelete      bool      `gorm:"column:is_delete"`
	JoinMethod    string    `gorm:"column:join_method"`
	Question      string    `gorm:"column:question"`
	Answer        string    `gorm:"column:answer"`
}

// TableName 组织表
func (Group) TableName() string {
	return "groups"
}

// Insert 创建
func (*Group) Create(group *Group) error {
	return db.Mysql.Insert(&group)
}

// GetGroupByID 根据ID获取组织
func (group *Group) GetGroupByID(groupID int) (*Group, error) {
	err := db.Mysql.Where("id = ?", groupID).First(group).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Info("没有此组织","groupId",gconv.String(groupID))
		}
		return nil, err
	}
	return group, nil
}

// GetGroupsByUserID 获取用户创建的组织
func (group *Group) GetUserCreateGroups(openID string) (*[]Group, error) {
	var groups []Group
	err := db.Mysql.Table("groups").Where("create_user = ?", openID).Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return &groups, nil
}
