package models

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"leastMall_gin/conn"
	"log"
	"time"
)

type cacheDb struct{}

var CacheDb = &cacheDb{}
var redisClient = conn.Redis

// Set 写入数据的方法
func (c cacheDb) Set(ctx *gin.Context, key string, value interface{}, redisTime int64) {
	bytes, _ := json.Marshal(value)
	redisClient.Set(ctx, key, string(bytes), time.Second*time.Duration(redisTime))
}

// Get 获取数据的方法
func (c cacheDb) Get(ctx *gin.Context, key string, obj interface{}) bool {
	redisStr, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		log.Printf("在redis里面读取数据失败：%v", err)
	}
	fmt.Println("在redis里面读取数据...")
	err = json.Unmarshal([]byte(redisStr), obj)
	if err != nil {
		return false
	}
	return true
}
