package redis

import (
	"ApscIM/pkg/model/base/user"
	"github.com/redis/go-redis/v9"
)

type UserRds struct {
	rdb redis.UniversalClient
}

func NewUserRds() (user.RdsUserInterface, error) {
	return &UserRds{}, nil
}
