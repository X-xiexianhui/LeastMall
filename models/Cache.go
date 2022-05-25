package models

import (
	"encoding/json"
	"fmt"
	"leastMall_gin/conn"
	"time"
)

type cacheDb struct{}

var CacheDb = &cacheDb{}
var redisClient = conn.Redis

// Set 写入数据的方法
func (c cacheDb) Set(key string, value interface{}) {
	if enableRedis {
		bytes, _ := json.Marshal(value)
		redisClient.Put(key, string(bytes), time.Second*time.Duration(redisTime))
	}
}

// Get 获取数据的方法
func (c cacheDb) Get(key string, obj interface{}) bool {
	if enableRedis {
		if redisStr := redisClient.Get(key); redisStr != nil {
			fmt.Println("在redis里面读取数据...")
			redisValue, ok := redisStr.([]uint8)
			if !ok {
				fmt.Println("获取redis数据失败")
				return false
			}
			json.Unmarshal([]byte(redisValue), obj)
			return true
		}
		return false
	}
	return false
}
