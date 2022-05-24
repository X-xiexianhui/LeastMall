//Package leastMall_gin
/*
   @author:xie
   @date:2022/5/23
   @note:
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"leastMall_gin/modules"
	"log"
)

var Conf modules.Config

func init() {
	cfg := viper.New()
	cfg.SetConfigName("config")
	cfg.SetConfigFile("./conf/app.yaml")
	if err := cfg.ReadInConfig(); err != nil { // 必须 先 读取 `ReadInConfig`
		log.Panicln(err)
	}
	err := cfg.Unmarshal(&Conf)
	fmt.Println(Conf)
	if err != nil {
		log.Panicln("参数配置失败")
	}
	cfg.WatchConfig()
	log.Println("参数配置成功")
}
func main() {
	modules.Db.LogMode(true)
	//日志着色
	gin.ForceConsoleColor()
	//启动gin
	r := gin.Default()
	modules.InitDataBase(Conf)
	_ = r.Run(":8080")
}
