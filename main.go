// Code generated by hertz generator.

package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app/server"
	"log"
	"sync"
	"time"

	"GraduateThesis/biz/service"
	"GraduateThesis/conf"
	"GraduateThesis/pkg/utils/locker"
	"GraduateThesis/router"
)

func main() {
	conf.Init()
	leaderCond := sync.NewCond(&sync.Mutex{})
	slaveCond := sync.NewCond(&sync.Mutex{})
	c, cancel := context.WithCancel(context.Background())
	service.Init(c, leaderCond, slaveCond)
	log.SetFlags(log.Llongfile | log.Ldate)

	go func() {
		lock := locker.NewLocker(locker.Config{
			Type: "redis",
			Redis: locker.RedisConfig{
				Addr:     conf.GConfig.RedisAddr,
				Password: conf.GConfig.RedisPassword,
				DB:       conf.GConfig.RedisDB,
			},
		})

		key := conf.GConfig.RedHaLeader
		value := time.Now().UnixNano()
		ticker := time.NewTicker(10 * time.Second)

		defer func() {
			lock.Unlock(context.TODO(), key, value)
			lock.Close()
			ticker.Stop()
		}()

		for {
			select {
			case <-c.Done():
				return
			case <-ticker.C:
				result, err := lock.Lock(context.TODO(), key, value, 1000*30)
				if err != nil {
					log.Println("抢锁报错")
				}
				if result {
					log.Println("抢锁成功，成为leader")
					leaderCond.Broadcast()
				} else {
					log.Println("抢锁失败，成为slave")
					slaveCond.Broadcast()
				}
			}
		}

	}()

	h := server.Default()

	router.Register(h)
	h.Spin()

	cancel()
}
