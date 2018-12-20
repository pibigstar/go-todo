package db

import (
	"github.com/jinzhu/gorm"
	"github.com/pibigstar/go-todo/config"
	"github.com/pibigstar/go-todo/utils/logger"
)

func init() {
	dbCfg := config.GetDBConfig()
	mysql := &mysql{}
	regist(mysql)
	createConnection(dbCfg)
}

var log = logger.New("db")

// myDB 封装一些数据库常用的操作
type myDB struct {
	*gorm.DB
}

func (db *myDB) Create(value interface{}) error {
	return db.Model(value).Create(value).Error
}

func (db *myDB) FindOne(result interface{}, where ...interface{}) error {
	return db.First(result, where...).Error
}
