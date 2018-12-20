package config

import (
	"github.com/pibigstar/go-todo/utils/logger"
	"github.com/spf13/viper"
)

var log = logger.New("config")

var ServerConfig serverConfig

type serverConfig struct {
	Port   int64
	Appid  string
	Secret string
}

// LoadConfig 加载配置文件
func LoadConfig() {
	// 设置配置文件名
	viper.SetConfigName("config")
	// 设置配置文件路径
	viper.AddConfigPath("config")

	// 解析配置
	viper.ReadInConfig()
}

// GetDBConfig 获取db配置
func GetDBConfig() map[string]interface{} {
	return viper.GetStringMap("db")
}

// GetServerConfig 获取服务器配置
func GetServerConfig() map[string]interface{} {
	return viper.GetStringMap("server")
}

// GetServerConfig 获取服务器配置
func buildServerConfig() {
	cfg := GetServerConfig()
	port := cfg["port"].(int64)
	ServerConfig = serverConfig{
		Port:   port,
		Appid:  cfg["appid"].(string),
		Secret: cfg["secret"].(string),
	}
}

func init() {
	LoadConfig()
	buildServerConfig()
}
