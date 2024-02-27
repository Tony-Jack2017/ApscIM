package db

import (
	"ApscIM/pkg/common/config"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func RunRedis() {

	conf := config.Conf.Database.Redis

	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Password, // 没有密码，默认值
		DB:       0,             // 默认DB 0
	})
	fmt.Println(rdb)
}
