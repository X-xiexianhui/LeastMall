//Package conn
/*
   @author:xie
   @date:2022/5/24
   @note:
*/
package conn

import (
	"context"
	"github.com/go-redis/redis/v8"
	"leastMall_gin/models"
	"log"
)

var Redis *redis.Client

func init() {
	cache := models.Conf.Redis
	Redis = redis.NewClient(&redis.Options{
		Addr:        cache.Host + ":" + cache.Port,
		Password:    cache.Password,
		DB:          cache.DefaultDB,
		DialTimeout: cache.DialTimeout,
	})
	_, err := Redis.Ping(context.TODO()).Result()
	if err != nil {
		log.Panic(err)
	}
	log.Println("redis连接成功……")
}
