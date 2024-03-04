package db

import (
	"ApscIM/pkg/common/config"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type CacheRedis struct {
	client *redis.Client
	keys   []string
}

func NewRedisService() *CacheRedis {
	conf := config.Conf.Database.Redis
	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Password, // 没有密码，默认值
		DB:       0,             // 默认DB 0
	})
	return &CacheRedis{
		client: rdb,
	}
}

func (r *CacheRedis) CreateKeys(keys ...string) {
}

func (r *CacheRedis) DelKey(key string) error {
	return nil
}

func (r *CacheRedis) GetValueByKey(key string) {
}

func (r *CacheRedis) SetKeysValue() {
}
