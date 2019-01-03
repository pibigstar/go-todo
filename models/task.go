package models

import (
	"github.com/pibigstar/go-todo/models/db"
	"github.com/spf13/cast"
	"time"
)

var MTask = &Task{}

type Task struct {
	ID             int
	TaskTitle      string    `gorm:"column:task_title"`
	TaskContent    string    `gorm:"column:task_content"`
	AppointTo      string    `gorm:"column:appoint_to"`
	CreateUser     string    `gorm:"column:create_user"`
	GroupID        int       `gorm:"column:group_id"`
	Status         int       `gorm:"column:status"`
	Tips           string    `gorm:"column:tips"`
	IsDelete       bool      `gorm:"column:is_delete"`
	IsRemind       bool      `gorm:"column:is_remind"`
	CompletionTime time.Time `gorm:"column:completion_time"`
	CreateTime     time.Time `gorm:"column:create_time"`
}

func (*Task) Name() string {
	return "task"
}

func (*Task) Create(task *Task) int{
	value, _ := db.Mysql.Create(task).Get("id")
	return cast.ToInt(value)
}
