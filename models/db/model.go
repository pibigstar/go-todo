package db

import (
	"github.com/spf13/cast"
)

// IDB 数据库模型层接口
type IDB interface {
	Name() string
	Init(conf map[string]interface{}) error
	Close()
}

var dbs []IDB

// Regist 加入到数组
func regist(db IDB) {
	dbs = append(dbs, db)
}

// createConnection 初始化数据库连接
func createConnection(conf map[string]interface{}) {
	for _, db := range dbs {
		if cfg, ok := conf[db.Name()]; ok {
			dbCfgMap := cast.ToStringMap(cfg)
			if err := db.Init(dbCfgMap); err != nil {
				log.Error("链接数据库错误", "db", db.Name(), "err", err.Error())
			} else {
				log.Info("链接据库成功", "db", db.Name())
			}
		}
	}
}
