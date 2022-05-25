package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

//定义结构体  缓存结构体 私有
type cookie struct{}

// Set 写入数据的方法
func (c cookie) Set(ctx *gin.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	ctx.SetCookie(key, string(bytes), 3600*24*30, "/", Conf.Domain, false, true)

}

// Remove 删除数据的方法
func (c cookie) Remove(ctx *gin.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	ctx.SetCookie(key, string(bytes), -1, "/", Conf.Domain, false, true)

}

// Get 获取数据的方法
func (c cookie) Get(ctx *gin.Context, key string, obj interface{}) bool {
	tempData, err := ctx.Request.Cookie(key)
	if err != nil {
		return false
	}
	err = json.Unmarshal([]byte(tempData.Value), obj)
	if err != nil {
		return false
	}
	return true

}

//实例化结构体
var Cookie = &cookie{}
