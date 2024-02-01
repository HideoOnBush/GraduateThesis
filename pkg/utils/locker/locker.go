package locker

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Type  string
	Redis RedisConfig
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type Locker interface {
	Lock(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	Unlock(ctx context.Context, key string, value interface{}) (bool, error)
	Renew(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)
	RenewWithPoll(ctx context.Context, key string, value interface{}, expiration time.Duration, interval time.Duration, tag string)
	Close() error
}

func NewLocker(lockerConfig Config) Locker {
	switch lockerConfig.Type {
	case "redis":
		redisClient := redis.NewClient(&redis.Options{
			Addr:     lockerConfig.Redis.Addr,
			Password: lockerConfig.Redis.Password,
			DB:       lockerConfig.Redis.DB,
		})
		return NewRedisLocker(redisClient)
	default:
		redisClient := redis.NewClient(&redis.Options{
			Addr:     lockerConfig.Redis.Addr,
			Password: lockerConfig.Redis.Password,
			DB:       lockerConfig.Redis.DB,
		})
		return NewRedisLocker(redisClient)
	}
}
