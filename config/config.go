package config

import (
	"github.com/pibigstar/go-todo/utils/logger"
	"github.com/spf13/viper"
)

var log = logger.New("config")

type common struct {
	Appid  string
	Secret string
}

// LoadConfig 加载配置文件
func LoadConfig() map[string]interface{} {
	// 设置配置文件名
	viper.SetConfigName("config")
	// 设置配置文件路径
	viper.AddConfigPath("config")
	viper.AddConfigPath("../../config")
	// 解析配置
	viper.ReadInConfig()
	// 获取db配置
	db := viper.GetStringMap("db")
	return db
}
