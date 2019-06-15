package models

import (
	"github.com/pibigstar/go-todo/models/db"
	"time"
)

// MAdmin 引用
var MAdmin = &Admin{}

// User 用户表
type Admin struct {
	ID         int       `gorm:"column:id"`
	UserName   string    `gorm:"column:username"`
	Desc       string    `gorm:"column:desc"`
	Salt       string    `gorm:"column:salt"`
	Password   string    `gorm:"column:password"`
	IsDelete   bool      `gorm:"column:is_delete"`
	CreateTime time.Time `gorm:"column:create_time"`
}

// TableName 用户表
func (*Admin) TableName() string {
	return "sys_admin"
}

func (t *Admin) Login(username, password string) (*Admin, error) {
	var admin Admin
	err := db.Mysql.Table(t.TableName()).Where("username = ? and password = ?", username, password).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
func (t *Admin) ListAdmin() ([]Admin, error) {
	var admins []Admin
	err := db.Mysql.Table(t.TableName()).Where("is_delete = ?", false).Find(&admins).Error
	if err != nil {
		return nil, err
	}
	return admins, nil
}
func (t *Admin) AdminDelete(id int) error {
	user := User{
		ID: id,
	}
	err := db.Mysql.Table(t.TableName()).Delete(&user).Error
	return err
}
