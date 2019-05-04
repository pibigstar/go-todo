package db

import (
	"fmt"
	"time"

	goRedis "github.com/go-redis/redis"
	"github.com/spf13/cast"
)

var Redis *redisClient

type redisClient struct {
	*goRedis.Client
}

type redis struct {
	client *redisClient
}

func (*redis) Name() string {
	return "redis"
}

func (*redis) Init(conf map[string]interface{}) error {
	client := goRedis.NewClient(&goRedis.Options{
		Addr: fmt.Sprintf("%s:%d", cast.ToString(conf["host"]), cast.ToInt(conf["port"])),
		Password: cast.ToString(conf["password"]),
		DB:   cast.ToInt(conf["db"]),
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Error("redis 链接失败", "err", err.Error())
	}
	Redis = &redisClient{client}
	return nil
}

func (redis *redis) Close() {
	redis.Close()
}

// 封装一些常用操作
func (client *redisClient) RSet(key string, value interface{}, expiration time.Duration) error {
	_, err := client.SetNX(key, value, expiration).Result()
	if err != nil {
		return err
	}
	return nil
}

func (client *redisClient) RGet(key string) interface{} {
	result, err := client.Get(key).Result()
	if err == goRedis.Nil {
		log.Info("key 值不存在", "key", key)
		return nil
	}
	return result
}
