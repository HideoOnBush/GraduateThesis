package locker

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type RedisLocker struct {
	client *redis.Client
}

func NewRedisLocker(client *redis.Client) Locker {
	return RedisLocker{client: client}
}

func (r RedisLocker) Lock(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	lua := `if redis.call("GET", KEYS[1]) == ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
    return "OK"
else
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
end`
	val, err := r.client.Eval(ctx, lua, []string{key}, value, expiration).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		log.SetFlags(log.Llongfile | log.Ldate)
		log.Printf("获取锁时报错, k=v: %s=%v, err: %s\n", key, value, err.Error())
		return false, err
	} else if val == nil {
		return false, nil
	}
	reply, ok := val.(string)
	if ok && reply == "OK" {
		return true, nil
	} else {
		log.Printf("Unknown reply when acquiring lock for %s: %v\n", key, reply)
		return false, nil
	}
}

func (r RedisLocker) Unlock(ctx context.Context, key string, value interface{}) (bool, error) {
	lua := `
-- 如果当前值与锁值一致,删除key
if redis.call('GET', KEYS[1]) == ARGV[1] then
	return redis.call('DEL', KEYS[1])
else
	return 0
end
`
	scriptKeys := []string{key}
	val, err := r.client.Eval(ctx, lua, scriptKeys, value).Result()
	return val == int64(1), err
}

func (r RedisLocker) Renew(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	lua := `
-- 如果当前值与锁值一致,删除key
if redis.call('GET', KEYS[1]) == ARGV[1] then
	return redis.call('PEXPIRE', KEYS[1], ARGV[2])
else
	return 0
end
`
	scriptKeys := []string{key}
	val, err := r.client.Eval(ctx, lua, scriptKeys, value, expiration/time.Millisecond).Result()
	return val == int64(1), err
}

func (r RedisLocker) RenewWithPoll(ctx context.Context, key string, value interface{}, expiration time.Duration, interval time.Duration, tag string) {
	for {
		t := time.After(interval)
		select {
		case <-ctx.Done():
			log.Printf("%s任务完成，关闭%s=%v的自动续期\n", tag, key, value)
		case <-t:
			result, err := r.Renew(ctx, key, value, expiration)
			if err != nil || !result {
				log.Printf("%s任务未完成，%s=%v自动续期失败，err: %v\n", tag, key, value, err)
			}
		}
	}
}

func (r RedisLocker) Close() error {
	return r.client.Close()
}
