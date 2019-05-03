package models

import (
	"github.com/pibigstar/go-todo/config"
	"github.com/pibigstar/go-todo/models/db"
	"github.com/pibigstar/go-todo/utils/logger"
	myLog "log"
)

var log = logger.New("models")

type MYSQLLogger struct {
}

func (logger *MYSQLLogger) Print(values ...interface{}) {
	var (
		level           = values[0]
		source          = values[1]
	)
	if level == "sql" {
		sql := values[3].(string)
		myLog.Println(sql, level, source)
	} else {
		myLog.Println(values)
	}
}


func init() {
	logger := &MYSQLLogger{}
	if config.ServerConfig.ShowSQL {
		db.Mysql.LogMode(true)
		db.Mysql.SetLogger(logger)
	}
}
