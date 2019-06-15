package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pibigstar/go-todo/models/db"
)

var MTask = &Task{}

type Task struct {
	ID             int       `gorm:"column:id;primary_key"`
	TaskTitle      string    `gorm:"column:task_title"`
	TaskContent    string    `gorm:"column:task_content"`
	TaskHtml       string    `gorm:"column:task_html"`
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
	FileIds        string    `gorm:"column:fileIds"`
}

func (*Task) TableName() string {
	return "task"
}

func (*Task) Create(task *Task) error {
	err := db.Mysql.Table(task.TableName()).Create(task).Error
	if err != nil {
		return err
	}
	return nil
}
func (task *Task) ListTask(openId string, status int, title string) ([]Task, error) {
	var tasks []Task
	model := db.Mysql.Table(task.TableName()).Where("appoint_to = ? and status = ?", openId, status)
	if title != "" {
		title = "%" + title + "%"
		model = model.Where("task_title like ?", title)
	}
	err := model.Order("create_time desc").Find(&tasks).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tasks, nil
}
func (task *Task) ChangeStatus(id int, status int) error {
	err := db.Mysql.Table(task.TableName()).Where("id = ?", id).UpdateColumn("status", status).Error
	return err
}
func (task *Task) GetTask(id int) (*Task, error) {
	var taskModel Task
	err := db.Mysql.Table(task.TableName()).Where("id = ?", id).Find(&taskModel).Error
	if err != nil {
		return nil, err
	}
	return &taskModel, nil
}

func (task *Task) SetRead(id int) error {
	err := db.Mysql.Table(task.TableName()).Where("id = ?", id).UpdateColumn("is_read", "1").Error
	return err
}
func (t *Task) CountTask(openId string, status int) (int, error) {
	var count int
	err := db.Mysql.Table(t.TableName()).Where("appoint_to = ? and status = ?", openId, status).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (t *Task) TaskList() (*[]Task, error) {
	var tasks []Task
	err := db.Mysql.Table(t.TableName()).Where("is_delete = ?", false).Find(&tasks).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tasks, nil
}
func (t *Task) TaskDelete(id int) error {
	task := Task{
		ID: id,
	}
	err := db.Mysql.Table(t.TableName()).Delete(&task).Error
	return err
}
