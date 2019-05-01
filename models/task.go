package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pibigstar/go-todo/models/db"
)

var MTask = &Task{}

type Task struct {
	ID             int
	TaskTitle      string    `gorm:"column:task_title"`
	TaskContent    string    `gorm:"column:task_content"`
	AppointTo      string    `gorm:"column:appoint_to"`
	CreateUser     string    `gorm:"column:create_user"`
	GroupName      string    `gorm:"column:group_name"`
	GroupID        int       `gorm:"column:group_id"`
	Status         int       `gorm:"column:status"`
	Tips           string    `gorm:"column:tips"`
	IsDelete       bool      `gorm:"column:is_delete"`
	IsRead         bool      `gorm:"column:is_read"`
	IsRemind       bool      `gorm:"column:is_remind"`
	CompletionTime time.Time `gorm:"column:completion_time"`
	CreateTime     time.Time `gorm:"column:create_time"`
}

func (*Task) Name() string {
	return "task"
}

func (*Task) Create(task *Task) error {
	err := db.Mysql.Table("task").Create(task).Error
	if err != nil {
		return err
	}
	return nil
}
func (task *Task) ListTask(openId string, status int, title string) ([]Task, error) {
	var tasks []Task
	model := db.Mysql.Table("task").Where("appoint_to = ? and status = ?", openId, status)
	if title != "" {
		title = "%" + title + "%"
		model = model.Where("task_title like ?", title)
	}
	err := model.Find(&tasks).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tasks, nil
}
func (task *Task) ChangeStatus(id int, status int) error {
	err := db.Mysql.Table("task").Where("id = ?", id).UpdateColumn("status", status).Error
	return err
}
