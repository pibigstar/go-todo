package models

import (
	"fmt"
	"github.com/pibigstar/go-todo/constant"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/pibigstar/go-todo/models/db"
)

// MGroupUser 实体外部引用
var MGroupUser = &GroupUser{}

// GroupUser 组织用户关联实体
type GroupUser struct {
	ID         int       `gorm:"column:id;primary_key"`
	GroupID    int       `gorm:"group_id"`
	GroupName  string    `gorm:"group_name"`
	UserID     string    `gorm:"user_id"`   //用户OpenID
	UserName   string    `gorm:"user_name"` //用户名
	CreateTime time.Time `gorm:"column:create_time"`
	IsDelete   bool      `gorm:"column:is_delete"`
	IsCreate   bool      `gorm:"column:is_create"`
}

func (*GroupUser) TableName() string {
	return "group_user"
}

// Insert 创建
func (*GroupUser) Create(groupUser *GroupUser) error {
	return db.Mysql.Insert(&groupUser)
}

// GetGroupsByUserOpenID 获取用户加入的群
func (user *GroupUser) GetUserJoinGroups(openID string) (*[]Group, error) {
	var groupIDs []int
	var groups []Group
	err := db.Mysql.Table(user.TableName()).Where("user_id = ?", openID).Pluck("group_id", &groupIDs).Error
	if err == gorm.ErrRecordNotFound || len(groupIDs) == 0 {
		return nil, errors.New("此用户没有加入任何群")
	}
	err = db.Mysql.Find(&groups, "id in(?) and is_delete = ?", groupIDs, false).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("没有找到群")
	}
	return &groups, nil
}

// GetUserOpenIDs 获取某个群的所有群成员的OpenID
func (user *GroupUser) GetUserOpenIDs(groupID int) ([]string, error) {
	var userOpenIds []string
	err := db.Mysql.Table(user.TableName()).
		Where("group_id in (?)", groupID).
		Pluck("user_id", &userOpenIds).Error
	if err != nil {
		return nil, err
	}
	return userOpenIds, nil
}

func (*GroupUser) GetFormIds(openIds []string) []string {
	var formIds []string
	for _, id := range openIds {
		formId, err := db.Redis.Get(fmt.Sprintf(constant.Redis_Prefix_Form_ID, id)).Result()
		if err == nil {
			formIds = append(formIds, formId)
		}
	}
	return formIds
}
func (user *GroupUser) IsExist(openId string, groupId int) (bool, error) {
	var result = &GroupUser{}
	err := db.Mysql.Table(user.TableName()).Where("user_id = ? and group_id = ?", openId, groupId).Find(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	if result != nil {
		return true, nil
	}
	return false, nil
}
func (user *GroupUser) GetUsers(groupId int) ([]GroupUser, error) {
	var users []GroupUser
	err := db.Mysql.Table(user.TableName()).Where("group_id = ?", groupId).Find(&users).Error
	if err != nil {
		log.Error("获取此群下的成员失败", "GroupId", groupId)
		return nil, err
	}
	return users, nil
}

func (t *GroupUser) ListMyCreateGroup(openId string) ([]GroupUser, error) {
	var groups []GroupUser
	err := db.Mysql.Table(t.TableName()).Where("user_id = ? and is_create = ?", openId, true).Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (t *GroupUser) ListMyJoinGroup(openId string) ([]GroupUser, error) {
	var groups []GroupUser
	err := db.Mysql.Table(t.TableName()).Where("user_id = ? and is_create = ?", openId, false).Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return groups, nil
}
