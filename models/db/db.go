package db

import (
	"github.com/pibigstar/go-todo/config"
	"github.com/pibigstar/go-todo/utils/logger"
	"github.com/spf13/cast"
)

var log = logger.New("db")

func init() {
	dbCfg := config.GetDBConfig()
	mysql := &mysql{}
	redis := &redis{}
	regist(mysql)
	regist(redis)
	createConnection(dbCfg)
}

// IDB 数据库模型层接口
type iDB interface {
	Name() string
	Init(conf map[string]interface{}) error
	Close()
}

var dbs []iDB

// Regist 加入到数组
func regist(db iDB) {
	dbs = append(dbs, db)
}

// createConnection 初始化数据库连接
func createConnection(conf map[string]interface{}) {
	for _, db := range dbs {
		if cfg, ok := conf[db.Name()]; ok {
			dbCfgMap := cast.ToStringMap(cfg)
			if err := db.Init(dbCfgMap); err != nil {
				log.Error("链接数据库错误", "db", db.Name(), "err", err.Error())
				panic(err)
			} else {
				log.Info("链接据库成功", "db", db.Name())
			}
		}
	}
}
