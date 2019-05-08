package config

import (
	"flag"
	"fmt"

	"github.com/spf13/cast"

	"github.com/pibigstar/go-todo/utils/logger"
	"github.com/spf13/viper"
)

var log = logger.New("config")

func init() {
	buildFlags()
	LoadConfig()
	buildServerConfig()
}

// ServerConfig 文件配置参数
var ServerConfig serverConfig

// ServerStartupFlags 启动自定义参数
var ServerStartupFlags serverStartupFlags

type serverConfig struct {
	Host            string
	Port            int
	Appid           string
	Secret          string
	WxLoginURL      string
	GroupCodeSecret string
	SecretKey       string
	ShowSQL         bool
}

type serverStartupFlags struct {
	Host        string
	Port        int
	Environment string
}

// LoadConfig 加载配置文件
func LoadConfig() {
	// 设置配置文件名
	configName := fmt.Sprintf("%s-%s", "config", ServerStartupFlags.Environment)
	viper.SetConfigName(configName)
	// 设置配置文件路径
	viper.AddConfigPath("conf")
	// 测试时使用路径
	viper.AddConfigPath("../../conf")
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

// buildServerConfig 构建文件服务器配置
func buildServerConfig() {
	cfg := GetServerConfig()
	ServerConfig = serverConfig{
		Port:            cast.ToInt(cfg["port"]),
		Appid:           cast.ToString(cfg["appid"]),
		Secret:          cast.ToString(cfg["secret"]),
		WxLoginURL:      cast.ToString(cfg["wxloginurl"]),
		GroupCodeSecret: cast.ToString(cfg["groupcodesecret"]),
		SecretKey:       cast.ToString(cfg["secretkey"]),
		ShowSQL:         cast.ToBool(cfg["showsql"]),
	}
	ServerConfig.Port = ServerStartupFlags.Port
	ServerConfig.Host = ServerStartupFlags.Host
}

// buildFlags 构建启动时参数配置
func buildFlags() {
	flag.StringVar(&ServerStartupFlags.Host, "host", "127.0.0.1", "listening host")
	flag.IntVar(&ServerStartupFlags.Port, "port", 7410, "listening port")
	flag.StringVar(&ServerStartupFlags.Environment, "env", "dev", "run time environment")
	if !flag.Parsed() {
		flag.Parse()
	}
}
