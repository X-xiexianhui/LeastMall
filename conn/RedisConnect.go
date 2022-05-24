//Package conn
/*
   @author:xie
   @date:2022/5/24
   @note:
*/
package conn

import (
	"github.com/go-redis/redis/v8"
	"leastMall_gin/modules"
	"log"
)

var Redis *redis.Client

func init() {
	cache := modules.Conf.Redis
	Redis = redis.NewClient(&redis.Options{
		Addr:        cache.Host + ":" + cache.Port,
		Password:    cache.Password,
		DB:          cache.DefaultDB,
		DialTimeout: cache.DialTimeout,
	})
	log.Println("redis连接成功……")
}
